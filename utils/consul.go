package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterServiceWithConsul(jsonData []byte, consulAddress string) error {
	jsonData, err := json.Marshal(jsonData)
	if err != nil {
		return fmt.Errorf("Error marshaling JSON: %v", err)
	}

	consulURL := fmt.Sprintf("http://%s/v1/agent/service/register", consulAddress)

	resp, err := http.Post(consulURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("Error sending request to Consul API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
