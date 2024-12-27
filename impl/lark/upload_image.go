package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yulai-123/message_send_tool/model"
	"io"
	"mime/multipart"
	"net/http"
)

var (
	oapiImageURL    = `https://open.feishu.cn/open-apis/im/v1/images`
	oapiBoundaryStr = `---7MA4YWxkTrZu0gW`
)

// uploadImage 飞书上传图片会使用 formData 格式上传，有点麻烦
// 1. 设置上传格式，包括 image, image_type
// 2. 发送请求
// 3. 解析返回结果
func (m *MessageHandler) uploadImage(imageInfo model.FileInfo) (string, error) {
	bu, err := m.parseImageParam(imageInfo)
	if err != nil {
		return "", err
	}
	token, err := m.getTenantAccessToken()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, oapiImageURL, bu)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+oapiBoundaryStr)
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := m.client.Do(req)
	if err != nil {
		return "", err
	}
	if resp == nil {
		err = fmt.Errorf("resp is nil")
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	messageResp := &oapiImageResponse{}
	err = json.Unmarshal(body, messageResp)
	if err != nil {
		return "", err
	}
	if messageResp.Code != 0 {
		err := fmt.Errorf("response error, %v, %v", messageResp.Code, messageResp.Msg)
		return "", err
	}

	return messageResp.Data.ImageKey, nil
}

// parseImageParam 将图片信息解析为 formData 格式
func (m *MessageHandler) parseImageParam(imageInfo model.FileInfo) (*bytes.Buffer, error) {
	bu := bytes.Buffer{}
	writer := multipart.NewWriter(&bu)
	err := writer.SetBoundary(oapiBoundaryStr)
	if err != nil {
		return nil, err
	}
	part, err := writer.CreateFormFile("image", imageInfo.FileName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, bytes.NewReader(imageInfo.FileData))
	if err != nil {
		return nil, err
	}
	err = writer.WriteField("image_type", "message")
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return &bu, nil
}
