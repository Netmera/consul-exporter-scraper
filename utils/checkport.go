package utils

import (
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Netmera/prometheus-consul-exporter/models"
	yaml "gopkg.in/yaml.v2"
)

// LoadConfigFromFile reads configuration from a YAML file
func LoadConfigFromFile(filename string) (*models.CheckPortModel, error) {
	filepath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config models.CheckPortModel
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func CheckPortOpen(port int) bool {
	address := net.JoinHostPort("localhost", strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
