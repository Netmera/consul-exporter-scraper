package models

type ServiceInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
	Meta    struct {
		Env  string `json:"env"`
		Type string `json:"type"`
	} `json:"meta"`
}
