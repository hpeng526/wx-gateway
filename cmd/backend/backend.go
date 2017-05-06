package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

const configFile = "./backend_config.json"

var configuration *ConfigFile
var wxServer *MqServer
var timeSec time.Duration

func init() {
	file, err := os.Open(configFile)
	decoder := json.NewDecoder(file)
	configuration = &ConfigFile{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Println("error:", err)
	}
	log.Println(configuration)
	wxServer = NewRedisMqServer(configuration.AppId, configuration.AppSecret, configuration.MqAddress)
	timeSec = time.Duration(configuration.delay) * time.Second
}

func main() {
	for {
		val := wxServer.Mq.Poll(configuration.Key, timeSec)
		wxServer.HandleMessage(val)
	}
}
