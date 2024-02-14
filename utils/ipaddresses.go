package utils

import (
	"net"

	"github.com/sirupsen/logrus"
)

// GetIPAddresses returns the list of IPv4 addresses of the machine
func GetIPAddresses() ([]net.IP, error) {
	var ips []net.IP

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logrus.Errorf("Error getting IP addresses: %v", err)
		return nil, err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP)
			}
		}
	}

	logrus.Infof("IP Addresses: %v", ips)
	return ips, nil
}
