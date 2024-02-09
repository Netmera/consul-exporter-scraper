package main

import (
	"fmt"
	"log"

	"github.com/Netmera/prometheus-consul-exporter/utils"
)

func main() {
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

	fmt.Println("Hostname:", hostname)
	fmt.Println("IP Addresses:")
	for _, ip := range ips {
		fmt.Println("-", ip)
	}

	// Load the configuration file
	config, err := utils.LoadConfigFromFile("exporter.yaml")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Check port for each exporter
	for _, exporter := range config.Exporters {
		fmt.Printf("Checking port for %s exporter...\n", exporter.Name)
		if utils.CheckPortOpen(exporter.Port) {
			fmt.Printf("Port %d is open for %s exporter\n", exporter.Port, exporter.Name)
		} else {
			fmt.Printf("Port %d is closed for %s exporter\n", exporter.Port, exporter.Name)
		}
	}
}
