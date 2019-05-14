package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
)

/* Structures */
type textRequest struct {
	Content string `json:"content"`
}

/* Private Methods */
func decodeStatus(res *http.Response) bool {
	status := struct {
		Status bool `json:"status"`
	}{}
	decoder := json.NewDecoder(res.Body)

	err := decoder.Decode(&status)

	if err != nil || !status.Status {
		return false
	}

	return true
}

/* Public Methods */
func TestServer() bool {
	// Test Text Server
	textServerTest := "http://" + TextServerPath + ":" + TextServerPort + "/hello"
	res, err := http.Get(textServerTest)

	if err != nil {
		ShowAlert("错误", "无法访问文字处理服务器")

		return false
	} else {
		flag := decodeStatus(res)

		if !flag {
			ShowAlert("错误", "文字处理服务器返回错误")
		}

		// Close
		_ = res.Body.Close()
	}

	// Test Image Server
	imageServerTest := "http://" + ImageServerPath + ":" + ImageServerPort + "/hello"
	res, err = http.Get(imageServerTest)

	if err != nil {
		ShowAlert("错误", "无法访问图片处理服务器")

		return false
	} else {
		flag := decodeStatus(res)

		if !flag {
			ShowAlert("错误", "图片处理服务器返回错误")
		}

		// Close
		_ = res.Body.Close()
	}

	return true
}

func PostTextServer(text string) bool {
	// Build JSON buffer
	req := textRequest{
		Content: text,
	}
	buff := new(bytes.Buffer)
	_ = json.NewEncoder(buff).Encode(req)

	// Send request
	textServerParse := "http://" + TextServerPath + ":" + TextServerPort + "/text"
	res, err := http.Post(textServerParse, "application/json; charset=utf-8", buff)

	// Handle response
	if err != nil {
		ShowAlert("错误", "访问文字处理服务器出现问题")

		return false
	} else {
		flag := decodeStatus(res)

		if !flag {
			ShowAlert("错误", "文字处理服务器处理错误")
		}

		// Close
		_ = res.Body.Close()
	}

	return true
}
