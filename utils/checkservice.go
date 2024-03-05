package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Netmera/consul-exporter-scraper/models"
	"github.com/sirupsen/logrus"
)

func CheckService(serviceName string, serviceAddress string, servicePort int, consulAddress string) bool {
	consulURL := fmt.Sprintf("http://%s:8500/v1/agent/services", serviceAddress)

	resp, err := http.Get(consulURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error making request to Consul: %v", resp.Status)
	}

	var services map[string]models.Service
	if err := json.NewDecoder(resp.Body).Decode(&services); err != nil {
		log.Fatal(err)
	}

	for _, service := range services {
		logrus.Infof("Service: %v", service)
		if service.Address == serviceAddress && service.Port == servicePort {
			return true
		}
	}
	return false
}
