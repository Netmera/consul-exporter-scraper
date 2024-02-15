package utils

import (
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Netmera/consul-exporter-scraper/models"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

// LoadConfigFromFile reads configuration from a YAML file
func LoadConfigFromFile(filename string) (*models.CheckPortModel, error) {
	logrus.Infof("Loading configuration from file: %s", filename)
	filepath, err := filepath.Abs(filename)
	if err != nil {
		logrus.Errorf("Error getting absolute path for file %s: %v", filename, err)
		return nil, err
	}

	file, err := os.Open(filepath)
	if err != nil {
		logrus.Errorf("Error opening file %s: %v", filepath, err)
		return nil, err
	}
	defer file.Close()

	var config models.CheckPortModel
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		logrus.Errorf("Error decoding YAML file %s: %v", filepath, err)
		return nil, err
	}

	logrus.Infof("Configuration loaded successfully from file: %s", filename)
	return &config, nil
}

// CheckPortOpen checks if the given port is open
func CheckPortOpen(port int) bool {
	logrus.Infof("Checking if port %d is open", port)
	address := net.JoinHostPort("localhost", strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		logrus.Errorf("Error checking port %d: %v", port, err)
		return false
	}
	defer conn.Close()
	logrus.Infof("Port %d is open", port)
	return true
}
