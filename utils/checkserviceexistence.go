package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func checkServiceExistence(serviceName string, consulAddress string) bool {
	consulURL := fmt.Sprintf("http://%s:8500/v1/agent/services", consulAddress)

	resp, err := http.Get(consulURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error making request to Consul: %v", resp.Status)
	}

	services := make(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(&services); err != nil {
		log.Fatal(err)
	}

	_, ok := services[serviceName]
	return ok
}
