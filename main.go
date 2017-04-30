package main

import (
	"errors"
	"io"
	"log"
	"net/http"

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
		for {
			w, err := conn.NextWriter(1)
			if err != nil {
				log.Println("unable to get write stream")
				return
			}
			messageType, r, err := conn.NextReader()
			if err != nil {
				log.Println("recive message fail, err:", err)
				break
			}
			switch messageType {
			case 1:
				log.Println("recived a text msg")
				io.Copy(w, r)
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
				conn.Close()
				break
			}
		}
	})

	e.Run("localhost:8080")
}
