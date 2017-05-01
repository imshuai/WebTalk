package main

import (
	"errors"
	"log"
	"net/http"
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
		conn, err := NewWebSocketUpgrader(1024, 1024).Upgrade(c.Writer, c.Request, nil)
		defer conn.Close()
		if err != nil {
			log.Println("faild to upgrade request to websocket")
			c.AbortWithError(http.StatusInternalServerError, errors.New("faild to upgrade request to websocket"))
			return
		}
		conn.SetCloseHandler(func(code int, text string) error {
			log.Println("user closed connection with code[", code, "] and text[", text, "]")
			return nil
		})
		conn.SetPingHandler(func(appData string) error {
			conn.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(time.Second))
			return nil
		})
		conn.SetPongHandler(func(appData string) error {
			conn.WriteControl(websocket.PingMessage, []byte(appData), time.Now().Add(time.Second))
			return nil
		})
		isAlive := make(chan bool)
		go func(conn *websocket.Conn, isAlive chan bool) {
			t := time.NewTicker(time.Second * 10)
			for {
				select {
				case <-t.C:
					select {
					case <-isAlive: //这里判断chan中是否有数据，保证之后读取的时候是空的
					case <-time.NewTicker(time.Millisecond * 100).C:
					}
					conn.WriteMessage(websocket.TextMessage, []byte("control:ping"))
					select {
					case i := <-isAlive:
						if i {
							log.Println("user is still alive, keep connection")
						} else {
							log.Println("user is not alive, closing the connection")
							conn.Close()
							return
						}
					case <-t.C:
						log.Println("user is not alive, closing the connection")
						conn.Close()
						return
					}
				}
			}
		}(conn, isAlive)
		for {
			messageType, r, err := conn.ReadMessage()
			if err != nil {
				log.Println("recive message fail, err:", err)
				break
			}
			switch messageType {
			case 1:
				if string(r) == "control:ping" {
					conn.WriteMessage(messageType, []byte("control:pong"))
				} else if string(r) == "control:pong" {
					isAlive <- true
				} else {
					log.Println("recived a text msg")
					conn.WriteMessage(messageType, r)
				}
			case 2:
				log.Println("recived a binary msg")
			case 8:
				log.Println("recived a control msg")
			case 9:
				log.Println("recived a ping msg")
			case 10:
				log.Println("recived a pong msg")
			default:
				log.Println("recived a unkown msg")
			}
		}
	})

	e.Run("localhost:8080")
}
