package lark

type oapiMessageRequest struct {
	ReceiveID string `json:"receive_id"`
	MsgType   string `json:"msg_type"`
	Content   string `json:"content"`
}

type oapiMessageResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type getTenantAccessTokenRequest struct {
	APPID     string `json:"app_id"`
	APPSecret string `json:"app_secret"`
}

type getTenantAccessTokenResponse struct {
	Code              int    `json:"code"`
	MSG               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int64  `json:"expire"`
}

type oapiImageRequest struct {
	ImageType string `json:"image_type"`
	Image     []byte `json:"image"`
}

type oapiImageResponse struct {
	Code int `json:"code"`
	Data struct {
		ImageKey string `json:"image_key"`
	} `json:"data"`
	Msg string `json:"msg"`
}
