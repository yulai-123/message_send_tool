package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type GetTenantAccessTokenRequest struct {
	APPID     string `json:"app_id"`
	APPSecret string `json:"app_secret"`
}

type GetTenantAccessTokenResponse struct {
	Code              int    `json:"code"`
	MSG               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int64  `json:"expire"`
}

func (l *LarkMessage) GetTenantAccessToken() (string, error) {
	if l.tenantAccessExpireTime.After(time.Now()) && l.tenantAccessToken != "" {
		return l.tenantAccessToken, nil
	}

	url := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"

	request := GetTenantAccessTokenRequest{
		APPID:     l.APPID,
		APPSecret: l.APPSecret,
	}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(requestBytes))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp == nil {
		err = fmt.Errorf("resp is nil")
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("request error, %v, %v", resp.StatusCode, resp.Status)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	tenantAccessTokenResp := &GetTenantAccessTokenResponse{}

	err = json.Unmarshal(body, tenantAccessTokenResp)
	if err != nil {
		return "", err
	}

	if tenantAccessTokenResp.Code != 0 {
		err := fmt.Errorf("response error, %v, %v", tenantAccessTokenResp.Code, tenantAccessTokenResp.MSG)
		return "", err
	}

	l.tenantAccessToken = tenantAccessTokenResp.TenantAccessToken
	l.tenantAccessExpireTime = time.Now().Add(time.Duration(tenantAccessTokenResp.Expire-1200) * time.Second)

	return l.tenantAccessToken, nil
}
