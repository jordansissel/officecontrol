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
	//avrQueue <- "PWON" // Power on
	//avrQueue <- "SICD" // Set Input to GAME
	//avrQueue <- "MV50" // Set master volume to 50

}
