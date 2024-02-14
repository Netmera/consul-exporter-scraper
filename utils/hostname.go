package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// GetHostname returns the hostname of the machine
func GetHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		logrus.Errorf("Error getting hostname: %v", err)
		return "", err
	}
	logrus.Infof("Hostname: %s", hostname)
	return hostname, nil
}
