package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yulai-123/message_send_tool/model"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	// 消息发送的url，这里默认用了openID
	oapiMessagesURL = `https://open.feishu.cn/open-apis/im/v1/messages?receive_id_type=open_id`

	// 用于解析消息发送的模版
	messageTemplate = `{"config": {"wide_screen_mode": true},"elements":[{"tag":"markdown","content":"{{.content}}"}{{range $k, $v := .imgList}},{"alt": {"content": "图片","tag": "plain_text"},"img_key": "{{$v}}","tag": "img"}{{end}}],"header": {"template": "blue","title": {"content": "{{.title}}","tag": "plain_text"}},"card_link": {"url": "{{.url}}","pc_url": "","android_url": "","ios_url": ""}}`
)

func (m *MessageHandler) sendMessage(message model.Message) error {
	content, err := m.parseMessage(message)
	if err != nil {
		return err
	}
	request := oapiMessageRequest{
		ReceiveID: message.ReceiveID,
		MsgType:   "interactive",
		Content:   content,
	}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}
	token, err := m.getTenantAccessToken()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, oapiMessagesURL, bytes.NewReader(requestBytes))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}
	if resp == nil {
		err = fmt.Errorf("resp is nil")
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	messageResp := &oapiMessageResponse{}
	err = json.Unmarshal(body, messageResp)
	if err != nil {
		return err
	}
	if messageResp.Code != 0 {
		err := fmt.Errorf("response error, %v, %v", messageResp.Code, messageResp.Msg)
		return err
	}

	return nil
}

func (m *MessageHandler) parseMessage(message model.Message) (string, error) {
	tmpl, err := template.New("basic").Parse(messageTemplate)
	if err != nil {
		return "", err
	}

	params := make(map[string]interface{})
	params["content"] = message.Content
	params["title"] = message.Title
	//todo params[] 这里需要补充图片的上传
	params["url"] = message.URL

	buffer := &bytes.Buffer{}
	err = tmpl.Execute(buffer, params)
	if err != nil {
		return "", err
	}

	result := buffer.String()

	result = strings.ReplaceAll(result, `\`, `\\`)
	result = strings.ReplaceAll(result, "\n", "\\n")

	return result, nil
}
