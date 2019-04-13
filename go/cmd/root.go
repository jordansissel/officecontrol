package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger
var logger *zap.SugaredLogger

var rootCmd = &cobra.Command{
	Use: "office",
}

func init() {
	zapLogger = zap.NewExample()
	logger = zapLogger.Sugar()
}

// Execute the top-level command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Errorw(err.Error())
		os.Exit(1)
	}
}
