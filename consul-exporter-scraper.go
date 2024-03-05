package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Netmera/consul-exporter-scraper/models"
	"github.com/Netmera/consul-exporter-scraper/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	// Check if running in Kubernetes environment
	inKubernetes := os.Getenv("KUBERNETES_SERVICE_HOST")

	if inKubernetes != "" {
		environment := os.Getenv("ENVIRONMENT")
		prometheusNameSpace := os.Getenv("PROMETHEUS_NAMESPACE")
		consulAddress := os.Getenv("CONSUL_ADDRESS")
		addresses := strings.Split(consulAddress, ",")
		for _, address := range addresses {
			logrus.Info(address)
		}
		// Get the node port of Prometheus service
		fmt.Println("Prometheus Namespace: ", prometheusNameSpace)
		nodePort, err := utils.GetPrometheusNodePort(prometheusNameSpace)
		if err != nil {
			logrus.Fatalf("Error getting Prometheus node port: %v", err)
		}

		// Get the hostnames of Kubernetes master nodes
		masterNodes, err := utils.GetMasterNodes()
		if err != nil {
			logrus.Fatalf("Error getting Kubernetes master node hostnames: %v", err)
		}

		for _, masterNode := range masterNodes {
			logrus.Info("Master Node: ", masterNode.Hostname)
			logrus.Info("Master Node IP: ", masterNode.IP)
			serviceInfo := models.ServiceInfo{
				ID:      masterNode.Hostname,
				Name:    "k8s-" + environment,
				Address: masterNode.IP,
				Port:    int(nodePort),
				Meta: struct {
					Env  string `json:"env"`
					Type string `json:"type"`
				}{Env: environment, Type: "kubernetes"},
			}

			// Convert struct to JSON
			jsonData, err := json.Marshal(serviceInfo)
			if err != nil {
				logrus.Fatalf("Error marshaling JSON: %v", err)
			}

			// Register service with Consul
			for _, consulAddress := range addresses {
				logrus.Info("Consul Address: ", consulAddress)
				consulURL := fmt.Sprintf("http://%s:8500/v1/agent/service/register", consulAddress)

				err = utils.RegisterServiceWithConsul(jsonData, consulURL)
				if err != nil {
					logrus.Warnf("Error registering service with Consul at %s: %v", consulAddress, err)
					continue
				}

				logrus.Infof("Service registered with Consul at %s", consulAddress)
				break
			}
			if err != nil {
				logrus.Fatalf("Failed to register service with any Consul addresses: %v", err)
			}
		}

	} else {
		environment := flag.String("environment", "", "Virtual Machine Environment")
		flag.Parse()
		// Get hostname
		hostname, err := utils.GetHostname()
		if err != nil {
			logrus.Fatalf("Error getting hostname: %v", err)
		}

		// Get IP addresses
		ips, err := utils.GetIPAddresses()
		if err != nil {
			logrus.Fatalf("Error getting IP addresses: %v", err)
		}

		// Load the configuration file
		config, err := utils.LoadConfigFromFile("/etc/consul-exporter-scraper/exporter.yaml")
		if err != nil {
			logrus.Fatalf("Error loading configuration: %v", err)
		}

		openPorts := make([]models.ExporterModel, 0)
		for _, exporter := range config.Exporters {
			if utils.CheckPortOpen(exporter.Port, ips[0].String()) {
				openPorts = append(openPorts, exporter)
			}
		}
		// Prepare data for Consul API
		for _, port := range openPorts {
			// Prepare data
			serviceInfo := models.ServiceInfo{
				ID:      hostname + "-" + port.ExportType,
				Name:    *environment + "-" + port.ExportType,
				Address: ips[0].String(),
				Port:    port.Port,
				Meta: struct {
					Env  string `json:"env"`
					Type string `json:"type"`
				}{Env: *environment, Type: port.ExportType},
			}

			fmt.Println("Service Info: ", serviceInfo)
			// Convert struct to JSON
			jsonData, err := json.Marshal(serviceInfo)
			if err != nil {
				logrus.Fatalf("Error marshaling JSON: %v", err)
			}

			// Register service with Consul
			for _, consulAddress := range config.ConsulAddresses {
				consulURL := fmt.Sprintf("http://%s/v1/agent/service/register", consulAddress)

				if utils.CheckService(serviceInfo.Name, consulAddress) {
					err = utils.RegisterServiceWithConsul(jsonData, consulURL)
					if err != nil {
						logrus.Warnf("Error registering service with Consul at %s: %v", consulAddress, err)
						continue
					}

					logrus.Infof("Service registered with Consul at %s", consulAddress)
					break
				} else {
					logrus.Warnf("Service already registered with Consul at %s", consulAddress)
				}
			}

			if err != nil {
				logrus.Fatalf("Failed to register service with any Consul addresses: %v", err)
			}
		}
	}

}
