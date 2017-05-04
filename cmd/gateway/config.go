package main

type ConfigFile struct {
	ServerAddress string `json:"server_address"`
	Database      string `json:"database"`
	MqAddress     string `json:"mq_address"`
}
