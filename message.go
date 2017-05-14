package main

import (
	"encoding/json"
	"errors"
)

//Message 定义消息结构
type Message struct {
	Author    string `json:"author"`
	Content   string `json:"content"`
	SendTo    string `json:"sendto"`
	Timestamp string `json:"timestamp"`
}

//MessageHandler 定义消息接收回调
func MessageHandler(msg string) error {
	m := new(Message)
	err := json.Unmarshal([]byte(msg), m)
	if err != nil {
		LogError("unmarshal message fail with error:", err)
		return errors.New("unmarshal message fail with error: " + err.Error())
	}
	if !users.Users[m.SendTo].reciveMsg(m) {
		LogError("send message from", m.Author, "to", m.SendTo, "fail")
		return errors.New("send message from " + m.Author + " to " + m.SendTo + " fail")
	}
	return nil
}
