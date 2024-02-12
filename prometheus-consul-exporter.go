package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/Netmera/prometheus-consul-exporter/models"
	"github.com/Netmera/prometheus-consul-exporter/utils"
)

func main() {
	Environment := flag.String("Environment", "", "Vırtual Machine Environment")
	flag.Parse()

	// Get hostname
	hostname, err := utils.GetHostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}

	// Get IP addresses
	ips, err := utils.GetIPAddresses()
	if err != nil {
		fmt.Println("Error getting IP addresses:", err)
		return
	}

	for _, ip := range ips {
		fmt.Println("-", ip)
	}

	// Load the configuration file
	config, err := utils.LoadConfigFromFile("configs/exporter.yaml")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

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
			ID:      *Environment,
			Name:    hostname,
			Address: ips[0].String(),
			Port:    port.Port,
			Meta: struct {
				Env  string `json:"env"`
				Type string `json:"type"`
			}{Env: *Environment, Type: port.ExportType},
		}

		// Convert struct to JSON
		jsonData, err := json.Marshal(serviceInfo)
		if err != nil {
			log.Fatalf("Error marshaling JSON: %v", err)
		}

		err = utils.RegisterServiceWithConsul(jsonData)
		if err != nil {
			log.Fatalf("Error registering service with Consul: %v", err)
		}

	}

}
