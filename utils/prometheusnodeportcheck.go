package utils

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GetPrometheusNodePort(namespace string) (int32, error) {

	config, err := rest.InClusterConfig()
	if err != nil {
		return 0, fmt.Errorf("error getting in-cluster config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return 0, fmt.Errorf("error creating Kubernetes client: %v", err)
	}

	serviceName := "prometheus-server"

	service, err := clientset.CoreV1().Services(namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
	if err != nil {
		return 0, fmt.Errorf("error getting service %s in namespace %s: %v", serviceName, namespace, err)
	}

	var nodePort int32
	for _, port := range service.Spec.Ports {
		if port.Name == "http" {
			nodePort = port.NodePort
			break
		}
	}

	if nodePort == 0 {
		return 0, fmt.Errorf("node port not found for service %s in namespace %s", serviceName, namespace)
	}

	return nodePort, nil
}
