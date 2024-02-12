package utils

import (
	"bytes"
	"fmt"
	"net/http"
)

func RegisterServiceWithConsul(jsonData []byte, consulAddress string) error {
	client := &http.Client{}

	// Create a new PUT request
	req, err := http.NewRequest("PUT", consulAddress, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("Error creating PUT request: %v", err)
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error sending PUT request to Consul API: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
