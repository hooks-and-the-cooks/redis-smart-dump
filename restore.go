package main

import (
	"github.com/sirupsen/logrus"
	"os/exec"
)

func restoreKeys(fileName string, host string, port string, logger *logrus.Logger) {
	command := exec.Command("redis-cli", "-h", host, port, "-c", "--pipe", "<", fileName)
	err := command.Run()
	if err != nil {
		logger.WithFields(logrus.Fields{"error": err}).Fatal("Not able to run restoreKeys command! exiting!")
	} else {
		logger.Info("Restore command ran successfully!")
	}
}
