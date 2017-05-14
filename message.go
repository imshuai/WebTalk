package main

//Message 定义消息结构
type Message struct {
	Author    string `json:"author"`
	Content   string `json:"content"`
	SendTo    string `json:"sendto"`
	Timestamp int64  `json:"timestamp"`
}

//MessageHandler 定义消息接收回调
func MessageHandler(msg string) error {
	return nil
}
