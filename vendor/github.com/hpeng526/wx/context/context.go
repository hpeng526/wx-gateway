package context

import "github.com/hpeng526/wx/cache"

type Context struct {
	AppID       string
	AppSecret   string
	AccessToken string
	Cache       cache.Cache
}
