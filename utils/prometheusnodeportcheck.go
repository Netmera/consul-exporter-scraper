package utils

import (
	"fmt"
	"os/exec"
)

// getPrometheusNodePort retrieves the node port of Prometheus service
func getPrometheusNodePort() (int, error) {
	cmd := exec.Command("kubectl", "get", "svc", "prometheus", "-n", "monitoring", "-o=jsonpath={.spec.ports[0].nodePort}")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("error running kubectl command: %v", err)
	}

	var nodePort int
	_, err = fmt.Sscanf(string(output), "%d", &nodePort)
	if err != nil {
		return 0, fmt.Errorf("error parsing node port: %v", err)
	}

	return nodePort, nil
}
