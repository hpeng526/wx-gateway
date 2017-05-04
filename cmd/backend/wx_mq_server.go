package main

import (
	"encoding/json"
	"fmt"
	"github.com/hpeng526/wx-gateway/mq"
	"github.com/hpeng526/wx/cache"
	"github.com/hpeng526/wx/context"
	"github.com/hpeng526/wx/template"
)

type MqServer struct {
	Ctx context.Context
	Mq  mq.MessageQueue
}

func NewRedisMqServer(appId string, appSecret string, serverAddr string) *MqServer {
	ctx := context.Context{AppID: appId, AppSecret: appSecret, Cache: cache.NewRedisCache(serverAddr)}
	redisMq := mq.NewRedisMq(serverAddr)
	return &MqServer{Ctx: ctx, Mq: redisMq}
}

func (ms *MqServer) HandleMessage(msg string) {
	if msg != "" {
		fmt.Printf("got msg: %s \n", msg)
		var tmd template.TemplateMessage
		json.Unmarshal([]byte(msg), &tmd)
		fmt.Printf("tmd is %v\n", tmd)
		tj, _ := json.Marshal(tmd)
		fmt.Printf("tj is %s\n", tj)
		token, err := ms.Ctx.GetAccessToken()
		if err != nil {
			fmt.Printf("accessToken err %s\n", token)
			return
		}
		tmd.SendTemplate(token)
	}

}
