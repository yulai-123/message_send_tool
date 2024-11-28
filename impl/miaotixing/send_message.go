package miaotixing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yulai-123/message_send_tool/model"
	"io/ioutil"
	"net/http"
)

var (
	triggerURL = "https://miaotixing.com/trigger"
)

type triggerRequest struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Type string `json:"type"`
}

type triggerResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (m *MessageHandler) sendMessage(message model.Message) error {
	content := message.Title + message.Content

	request := triggerRequest{
		ID:   m.triggerID,
		Text: content,
		Type: "json",
	}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, triggerURL, bytes.NewReader(requestBytes))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
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
	messageResp := &triggerResp{}
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
