package main

import (
	"encoding/json"
	"github.com/hpeng526/wx-gateway/mq"
	"github.com/hpeng526/wx-gateway/po"
	"github.com/hpeng526/wx-gateway/service"
	"github.com/hpeng526/wx/template"
	"log"
	"net/http"
	"os"
)

const configFile = "./gateway_config.json"

var configuration *ConfigFile
var redisMq *mq.RedisMq
var us *service.UserService

func userGateway(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t po.UserMessage
	err := decoder.Decode(&t)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := po.HttpResp{Success: false, Msg: "error param"}
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
		return
	}
	user, err := us.FindUserById(t.UserId)
	log.Printf("user is %v\n, err is %v\n", user, err)
	msg := template.TemplateMessage{
		ToUser:     user.UserWXId,
		TemplateID: user.TemplateId,
		URL:        t.URL,
		JSONData:   t.Data,
	}
	log.Printf("msg is %v", msg)

	if msg.ToUser != "" {
		// offer to mq
		jsonData, err := json.Marshal(msg)
		if err != nil {
			log.Printf("err %v", err)
		}
		go redisMq.Offer(configuration.Key, string(jsonData))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := po.HttpResp{Success: true, UserWxId: msg.ToUser}
	jsonResp, err := json.Marshal(resp)
	w.Write(jsonResp)

}

func init() {
	file, err := os.Open(configFile)
	decoder := json.NewDecoder(file)
	configuration = &ConfigFile{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Println("error:", err)
	}
	log.Println(configuration)
	redisMq = mq.NewRedisMq(configuration.MqAddress)
	us = &service.UserService{DataSource: configuration.Database}
}

func main() {
	http.HandleFunc("/u", userGateway)
	err := http.ListenAndServe(configuration.ServerAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
