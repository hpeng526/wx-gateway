package mq

import (
	"encoding/json"
	"fmt"
	"github.com/hpeng526/wx/template"
	"testing"
)

func TestTmpData(t *testing.T) {
	tmpStr := `{"touser":"oXymQwcLPXl8-nkJF6Z7bAbfCPGs","template_id":"TY33t4IkXbyobY_3PpKZu5h4LNxCoYNil9WyJAZOOZE","url":"http://z.cn","topcolor":"#ff0000","data":{"first":{"value":"first","color":"#173177"},"send":{"value":"send","color":"#173177"},"text":{"value":"Text","color":"#173177"},"time":{"value":"Time","color":"#173177"},"remark":{"value":"Remark","color":"#173177"}}}`

	var msg template.TemplateMessage

	json.Unmarshal([]byte(tmpStr), &msg)

	jsonData, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	fmt.Printf("tmp is : %v \n", string(jsonData))

	//redisMq := NewRedisMq("127.0.0.1:6379")
	//redisMq.Offer("testmq", string(jsonData))
}
