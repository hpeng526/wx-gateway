package main

type ConfigFile struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	MqAddress string `json:"mq_address"`
	Key       string `json:"key"`
	delay     int32  `json:"delay"`
}
