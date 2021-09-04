package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

var ctx = context.Background()

func main() {
	logger := logrus.New()
	app := createCLIApp(logger)

	err := app.Run(os.Args)
	if err != nil {
		logger.WithFields(logrus.Fields{"error": err}).Fatal("Not able to start the app! exiting!")
		os.Exit(1)
	}
}
