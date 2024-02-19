// utils/master_nodes.go

package utils

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/Netmera/consul-exporter-scraper/models"
)

func GetMasterNodes() ([]models.MasterNode, error) {
	cmd := exec.Command("kubectl", "get", "nodes", "--selector=node-role.kubernetes.io/master", "-o=json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error running kubectl command: %v", err)
	}

	var nodeList struct {
		Items []struct {
			Metadata struct {
				Name string `json:"name"`
			} `json:"metadata"`
			Status struct {
				Addresses []struct {
					Address string `json:"address"`
					Type    string `json:"type"`
				} `json:"addresses"`
			} `json:"status"`
		} `json:"items"`
	}
	if err := json.Unmarshal(output, &nodeList); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON output: %v", err)
	}

	var masterNodes []models.MasterNode
	for _, item := range nodeList.Items {
		for _, addr := range item.Status.Addresses {
			if addr.Type == "InternalIP" {
				masterNodes = append(masterNodes, models.MasterNode{
					Hostname: item.Metadata.Name,
					IP:       addr.Address,
				})
				break
			}
		}
	}

	return masterNodes, nil
}
