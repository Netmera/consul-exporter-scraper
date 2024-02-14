package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/Netmera/prometheus-consul-exporter/models"
	"github.com/Netmera/prometheus-consul-exporter/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	environment := flag.String("environment", "", "Virtual Machine Environment")
	flag.Parse()

	// Get hostname
	hostname, err := utils.GetHostname()
	if err != nil {
		logrus.Fatalf("Error getting hostname: %v", err)
	}

	// Get IP addresses
	ips, err := utils.GetIPAddresses()
	if err != nil {
		logrus.Fatalf("Error getting IP addresses: %v", err)
	}

	// Load the configuration file
	config, err := utils.LoadConfigFromFile("/etc/prometheus-consul-exporter/exporter.yaml")
	if err != nil {
		logrus.Fatalf("Error loading configuration: %v", err)
	}

	consulURL := fmt.Sprintf("http://%s/v1/agent/service/register", config.ConsulAddress)

	openPorts := make([]models.ExporterModel, 0)
	for _, exporter := range config.Exporters {
		if utils.CheckPortOpen(exporter.Port) {
			openPorts = append(openPorts, exporter)
		}
	}
	// Prepare data for Consul API
	for _, port := range openPorts {
		// Prepare data
		serviceInfo := models.ServiceInfo{
			ID:      hostname,
			Name:    *environment,
			Address: ips[0].String(),
			Port:    port.Port,
			Meta: struct {
				Env  string `json:"env"`
				Type string `json:"type"`
			}{Env: *environment, Type: port.ExportType},
		}

		// Convert struct to JSON
		jsonData, err := json.Marshal(serviceInfo)
		if err != nil {
			logrus.Fatalf("Error marshaling JSON: %v", err)
		}

		err = utils.RegisterServiceWithConsul(jsonData, consulURL)
		if err != nil {
			logrus.Fatalf("Error registering service with Consul: %v", err)
		}

		logrus.Infof("Service registered with Consul: %s", serviceInfo.Name)

	}

}
