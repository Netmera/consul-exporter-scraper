package utils

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/Netmera/consul-exporter-scraper/models"
)

func GetMasterNodes() ([]models.MasterNode, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("error getting in-cluster config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error creating Kubernetes client: %v", err)
	}

	labelSelector := "node-role.kubernetes.io/master"

	nodeList, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting master nodes: %v", err)
	}

	var masterNodes []models.MasterNode
	for _, node := range nodeList.Items {
		var internalIP string
		for _, addr := range node.Status.Addresses {
			if addr.Type == corev1.NodeInternalIP {
				internalIP = addr.Address
				break
			}
		}
		if internalIP != "" {
			masterNodes = append(masterNodes, models.MasterNode{
				Hostname: node.Name,
				IP:       internalIP,
			})
		}
	}

	return masterNodes, nil
}
