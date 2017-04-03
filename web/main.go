package main

import (
	"encoding/json"
	"fmt"
	"github.com/hpeng526/wx-backend/mq"
	"github.com/hpeng526/wx-gateway/po"
	"github.com/hpeng526/wx-gateway/service"
	"github.com/hpeng526/wx/template"
	"log"
	"net/http"
)

func gateway(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t po.UserMessage
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	us := service.UserService{DataSource: "./gateway.sqlite"}
	user, err := us.FindUserById(t.UserId)

	fmt.Printf("user is %s\n, err is %s\n", user, err)
	msg := template.TemplateMessage{
		ToUser:     user.UserWXId,
		TemplateID: user.TemplateId,
		URL:        t.URL,
		JSONData:   t.Data,
	}
	fmt.Printf("msg is %s", msg)

	if msg.ToUser != "" {
		// offer to mq
		redisMq := mq.NewRedisMq("127.0.0.1:6379")
		jsonData, err := json.Marshal(msg)
		if err != nil {
			fmt.Printf("err %v", err)
		}
		redisMq.Offer("testmq", string(jsonData))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := po.HttpResp{Success: true, UserWxId: msg.ToUser}
	jsonResp, err := json.Marshal(resp)
	w.Write(jsonResp)

}

func main() {
	http.HandleFunc("/u", gateway)           //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
