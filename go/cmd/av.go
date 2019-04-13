package cmd

import (
	"os"
	"time"

	"github.com/jordansissel/officecontrol/go/denon"
	"github.com/spf13/cobra"
)

var avCmd = &cobra.Command{
	Use:   "av",
	Run:   avRun,
	Short: "Control Denon A/V Reciever",
	Args:  cobra.ArbitraryArgs,
}

func init() {
	rootCmd.AddCommand(avCmd)
}

func avRun(cmd *cobra.Command, args []string) {
	avr := denon.Discover()

	if len(args) == 0 {
		status, err := avr.Query()
		if err != nil {
			logger.Errorw(err.Error())
			os.Exit(1)
		}

		logger.Infof("Power: %t", status.Power)
		logger.Infof("Volume: %.1f", status.Volume)
		logger.Infof("Input: %s", status.Input.Name)
		return
	}

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
