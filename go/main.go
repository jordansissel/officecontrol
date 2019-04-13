package main

import (
	"os"
	"time"

	"github.com/jordansissel/officecontrol/go/denon"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger
var logger *zap.SugaredLogger

func init() {
	zapLogger = zap.NewExample()
	logger = zapLogger.Sugar()
}

func main() {
	avr := denon.Discover()

	for i, command := range os.Args[1:] {
		avr.Command(command)
		if i+1 < len(os.Args[1:]) {
			time.Sleep(1 * time.Second)
		}
	}

	status, err := avr.Query()

	if err != nil {
		logger.Errorw("Failed to query Denon current status", "error", err)
		os.Exit(1)
	}

	logger.Infow("Status", "power", status.Power, "volume", status.Volume, "input", status.Input.Name)
}
