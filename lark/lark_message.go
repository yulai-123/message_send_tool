package lark

import (
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Timeout: 5 * time.Second,
	}
)

type LarkMessage struct {
	APPID                  string
	APPSecret              string
	tenantAccessToken      string
	tenantAccessExpireTime time.Time
}

func NewLarkMessage(appID, appSecret string) *LarkMessage {
	return &LarkMessage{
		APPID:                  appID,
		APPSecret:              appSecret,
		tenantAccessExpireTime: time.Now().Add(-1 * time.Hour),
	}
}
