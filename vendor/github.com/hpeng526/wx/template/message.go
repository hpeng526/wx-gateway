package template

import "encoding/json"

type TemplateMessage struct {
	ToUser     string          `json:"touser"`
	TemplateID string          `json:"template_id"`
	URL        string          `json:"url,omitempty"`
	TopColor   string          `json:"topcolor,omitempty"`
	JSONData   json.RawMessage `json:"data"`
}

type TemplateResp struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	MsgID   int64  `json:"msgid"`
}
