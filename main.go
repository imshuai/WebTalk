package main

import (
	"log"
	"time"

	"io"

	"github.com/imshuai/wshelper"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()
	e.LoadHTMLFiles("./tmpls/index.html")

	e.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	e.GET("/user/login/:nickname", func(c *gin.Context) {
		wsh := wshelper.NewWebSocketHelper(1024, 1024)
		wsh.KeepAlive = true
		wsh.CloseHandleFunc = func(code int, text string) error {
			log.Println("user closed connection with code[", code, "] and text[", text, "]")
			return nil
		}
		wsh.PingHandleFunc = func(pingMsg string) error {
			return wsh.WriteControl(wshelper.PingMessage, []byte(pingMsg), time.Now().Add(time.Second))
		}
		wsh.PongHandleFunc = func(pongMsg string) error {
			return wsh.WriteControl(wshelper.PingMessage, []byte(pongMsg), time.Now().Add(time.Second))
		}
		wsh.TextMsgHandleFunc = func(msg string) error {
			return wsh.WriteMessage(wshelper.TextMessage, msg)
		}
		wsh.StreamMsgHandleFunc = func(r io.Reader) error {
			pr, pw := io.Pipe()
			io.Copy(io.Writer(pw), r)
			return wsh.WriteMessage(wshelper.StreamMessage, io.Reader(pr))
		}
		wsh.StartHandle(c.Writer, c.Request)
	})

	e.Run("localhost:8080")
}
