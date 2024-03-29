package models

// CheckPortModel represents the structure of the configuration file
type CheckPortModel struct {
	ConsulAddresses []string        `yaml:"consuladdresses"`
	Exporters       []ExporterModel `json:"exporters"`
}

// ExporterModel represents exporter information
type ExporterModel struct {
	Name       string `json:"name"`
	Port       int    `json:"port"`
	ExportType string `json:"exporttype"`
}
