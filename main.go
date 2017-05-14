package main

import (
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

		wsh.StartHandle(c.Writer, c.Request)
	})

	e.Run("localhost:8080")
}
