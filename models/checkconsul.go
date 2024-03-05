package models

type Service struct {
	ID      string `json:"ID"`
	Address string `json:"Address"`
	Port    int    `json:"Port"`
}
