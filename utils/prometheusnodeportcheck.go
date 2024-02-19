package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetPrometheusNodePort retrieves the node port of Prometheus service
func GetPrometheusNodePort(namespace string) (int, error) {
	cmd := exec.Command("kubectl", "get", "svc", "prometheus-server", "-n", namespace, "-o=jsonpath={.spec.ports[0].nodePort}")
	fmt.Println("Command: ", cmd)
	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return 0, fmt.Errorf("error running kubectl command: %v, stderr: %s", err, string(exitError.Stderr))
		}
		return 0, fmt.Errorf("error running kubectl command: %v", err)
	}

	nodePortStr := strings.TrimSpace(string(output))
	var nodePort int
	_, err = fmt.Sscanf(nodePortStr, "%d", &nodePort)
	if err != nil {
		return 0, fmt.Errorf("error parsing node port: %v", err)
	}
	fmt.Println("Node Port: ", nodePort)
	return nodePort, nil
}
