package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

// getKubernetesMasterNodeHostnames retrieves the hostnames of Kubernetes master nodes
func getKubernetesMasterNodeHostnames() ([]string, error) {
	cmd := exec.Command("kubectl", "get", "nodes", "--selector=node-role.kubernetes.io/master", "-o=jsonpath={.items[*].metadata.name}")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error running kubectl command: %v", err)
	}

	hostnames := strings.Fields(string(output))
	return hostnames, nil
}
