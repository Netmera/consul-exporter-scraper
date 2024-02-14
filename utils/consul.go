package utils

import (
	"bytes"
	"net/http"

	"github.com/sirupsen/logrus"
)

func RegisterServiceWithConsul(jsonData []byte, consulAddress string) error {
	logrus.Infof("Registering service with Consul at address: %s", consulAddress)

	client := &http.Client{}

	// Create a new PUT request
	req, err := http.NewRequest("PUT", consulAddress, bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Errorf("Error creating PUT request to Consul API: %v", err)
		return nil
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("Error sending PUT request to Consul API: %v", err)
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("Unexpected status code: %d", resp.StatusCode)
		return nil
	}

	return nil
}
