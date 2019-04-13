package main

import (
	"github.com/jordansissel/officecontrol/go/cmd"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger
var logger *zap.SugaredLogger

func init() {
	zapLogger = zap.NewExample()
	logger = zapLogger.Sugar()
}

func main() {
	cmd.Execute()
}
