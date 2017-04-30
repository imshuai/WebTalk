package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

//WebSocketUpgrader 用来升级http协议到ws协议
type WebSocketUpgrader struct {
	websocket.Upgrader
}

//NewWebSocketUpgrader 创建新的请求升级器
//rbufSize 指定读缓存大小，字节
//wbufSize 指定写缓存大小，字节
func NewWebSocketUpgrader(rbufSize, wbufSize int) *WebSocketUpgrader {
	return &WebSocketUpgrader{
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  rbufSize,
			WriteBufferSize: wbufSize,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}
