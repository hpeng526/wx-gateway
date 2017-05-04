package context

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const AccessTokenURL = "https://api.weixin.qq.com/cgi-bin/token"

type RespAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (ctx *Context) GetAccessToken() (accessToken string, err error) {
	cacheKey := fmt.Sprintf("access_token_%s", ctx.AppID)
	value := ctx.Cache.Get(cacheKey)

	if value != nil {
		accessToken = value.(string)
		return
	}

	var respAccessToken RespAccessToken
	respAccessToken, err = ctx.GetAccessTokenFromServer()
	if err != nil {
		return
	}
	accessToken = respAccessToken.AccessToken
	return

}

func (ctx *Context) GetAccessTokenFromServer() (token RespAccessToken, err error) {
	url := fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", AccessTokenURL, ctx.AppID, ctx.AppSecret)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &token)
	if err != nil {
		return
	}
	expires := token.ExpiresIn - 1500
	accessTokenCacheKey := fmt.Sprintf("access_token_%s", ctx.AppID)
	err = ctx.Cache.Set(accessTokenCacheKey, token.AccessToken, time.Duration(expires)*time.Second)
	return
}
