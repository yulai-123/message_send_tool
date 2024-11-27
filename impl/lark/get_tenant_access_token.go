package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	// 用于获取 tenant_access_token
	oapiTenantAccessToken = `https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal`
)

func (m *MessageHandler) getTenantAccessToken() (string, error) {
	if m.tenantAccessExpireTime.After(time.Now()) && m.tenantAccessToken != "" {
		return m.tenantAccessToken, nil
	}

	request := getTenantAccessTokenRequest{
		APPID:     m.appID,
		APPSecret: m.appSecret,
	}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", oapiTenantAccessToken, bytes.NewReader(requestBytes))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := m.client.Do(req)
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

	tenantAccessTokenResp := &getTenantAccessTokenResponse{}
	err = json.Unmarshal(body, tenantAccessTokenResp)
	if err != nil {
		return "", err
	}

	if tenantAccessTokenResp.Code != 0 {
		err := fmt.Errorf("response error, %v, %v", tenantAccessTokenResp.Code, tenantAccessTokenResp.MSG)
		return "", err
	}

	m.tenantAccessToken = tenantAccessTokenResp.TenantAccessToken
	m.tenantAccessExpireTime = time.Now().Add(time.Duration(tenantAccessTokenResp.Expire-1200) * time.Second)

	return m.tenantAccessToken, nil
}
