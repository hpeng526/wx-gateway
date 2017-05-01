package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const SendTemplateURL = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"

func (tm *TemplateMessage) SendTemplate(token string) (tResp TemplateResp, err error) {
	url := fmt.Sprintf(SendTemplateURL, token)
	fmt.Printf("url is : %s\n", url)
	tmData, err := json.Marshal(tm)
	if err != nil {
		fmt.Print("error post")
	}
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewReader(tmData))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &tResp)
	fmt.Printf("body is : %v, tResp is %v \n", string(body), tResp)
	return
}
