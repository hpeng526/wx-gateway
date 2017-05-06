package main

type ConfigFile struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	MqAddress string `json:"mq_address"`
	Key       string `json:"key"`
	Delay     int32  `json:"delay"`
}
