package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()
	e.LoadHTMLFiles("./tmpls/index.html")

	e.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	e.GET("/ws", func(c *gin.Context) {
		wsh := NewWebSocketHelper(1024, 1024)
		wsh.CloseHandleFunc = func(code int, text string) error {
			log.Println("user closed connection with code[", code, "] and text[", text, "]")
			wsh.isAlive <- false
			return nil
		}
		wsh.PingHandleFunc = func(pingMsg string) error {
			wsh.WriteControl(websocket.PongMessage, []byte(pingMsg), time.Now().Add(time.Second))
			return nil
		}
		wsh.PongHandleFunc = func(pongMsg string) error {
			wsh.WriteControl(websocket.PingMessage, []byte(pongMsg), time.Now().Add(time.Second))
			return nil
		}
		wsh.MessageHandleFunc = func(messageType int, msg []byte, err error) error {
			return nil
		}
	})

	e.Run("localhost:8080")
}
