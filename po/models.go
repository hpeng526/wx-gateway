package po

import (
	"encoding/json"
	"time"
)

type User struct {
	UserId     int64
	UserWXId   string
	TemplateId string
	CreateTime time.Time
}

// post data
type UserMessage struct {
	UserId int64           `json:"user_id"`
	URL    string          `json:"url,omitempty"`
	Data   json.RawMessage `json:"data"`
}

type HttpResp struct {
	Success  bool   `json:"success"`
	UserWxId string `json:"user_wx_id,omitempty"`
}
