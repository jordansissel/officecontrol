package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"./denon"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger
var logger *zap.SugaredLogger

func init() {
	zapLogger = zap.NewExample()
	logger = zapLogger.Sugar()
}

func mqttClient() mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker("tcp://office:1883")
	opts.SetKeepAlive(5 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	return mqtt.NewClient(opts)
}

func watchDenon(avr *denon.Denon, client mqtt.Client) {
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		result, err := avr.Query()
		if err != nil {
			logger.Error("Failed querying Denon AVR")
		}
		logger.Debugw("Denon AVR Status", "volume", result.MasterVolume)
		client.Publish("/denon/status", 0, false, fmt.Sprintf("Volume: %.0f", result.MasterVolume))
	}
}

// Send commands to the receiver.
// This function introduces a short delay between commands to avoid softlocking.
// In practice, sending two commands (like PWON and SIGAME) too close together
// will softlock my AVR-720w requiring a hard power cycle (aka unplug power).
func sendDenon(avr *denon.Denon, c chan string) {
	for command := range c {
		logger.Infow("Sending Denon AVR command", "command", command)
		avr.Command(command)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	client := mqttClient()
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	avr := denon.Discover()
	go watchDenon(avr, client)

	avrQueue := make(chan string, 5)
	go sendDenon(avr, avrQueue)

	client.Subscribe("/denon/command", 0, func(client mqtt.Client, msg mqtt.Message) {
		command := string(msg.Payload())
		avrQueue <- command
	})

	client.Subscribe("/office/panel/online", 0, func(client mqtt.Client, msg mqtt.Message) {
		logger.Info("The panel is online.")
		// There's no message here. This is the signal the panel sends when it powers up and gets network.
		// This signal is a strong indicator that I just sat down and probably want some things to power on.

		args := []string{"power", "-N", "pork-ipmi", "-u", "-U", "root", "-P", os.ExpandEnv("${ipmipass}")}
		logger.Infow("Sending ipmiutil", "args", args)
		cmd := exec.Command("ipmiutil", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

		// Turn on the receiver
		avrQueue <- "PWON"   // Power on
		avrQueue <- "SIGAME" // Set Input to GAME
		avrQueue <- "MV50"   // Set master volume to 50
	})

	// Block forever.
	select {}
}
