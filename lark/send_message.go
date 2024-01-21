package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	messageTemplate = `{"config": {"wide_screen_mode": true},"elements":[{"tag":"markdown","content":"{{.content}}"}{{range $k, $v := .imgList}},{"alt": {"content": "图片","tag": "plain_text"},"img_key": "{{$v}}","tag": "img"}{{end}}],"header": {"template": "blue","title": {"content": "{{.title}}","tag": "plain_text"}}}`
)

func (l *LarkMessage) parseMessage(data MessageData) (string, error) {
	tmpl, err := template.New("basic").Parse(messageTemplate)
	if err != nil {
		return "", err
	}

	params := make(map[string]interface{})
	params["content"] = data.Content.Content
	params["title"] = data.Content.Title
	params["imgList"] = data.Content.ImgList

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

type SendMessageRequest struct {
	ReceiveID string `json:"receive_id"`
	MsgType   string `json:"msg_type"`
	Content   string `json:"content"`
}

type SendMessageResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (l *LarkMessage) SendMessage(messageData MessageData) error {
	url := "https://open.feishu.cn/open-apis/im/v1/messages?receive_id_type=open_id"

	content, err := l.parseMessage(messageData)
	if err != nil {
		return err
	}

	request := SendMessageRequest{
		ReceiveID: messageData.ReceiveID,
		MsgType:   "interactive",
		Content:   content,
	}

	requestBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}
	fmt.Println(string(requestBytes))

	req, err := http.NewRequest("POST", url, bytes.NewReader(requestBytes))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	token, err := l.GetTenantAccessToken()
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
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

	messageResp := &SendMessageResponse{}

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
