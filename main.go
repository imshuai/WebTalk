package main

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()
	e.LoadHTMLFiles("./tmpls/index.html")
	e.Static("/statics", "./statics")

	e.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	e.GET("/user/login/:nickname", func(c *gin.Context) {
		u := UserOnline(c.Param("nickname"))
		users.Add(u)
		u.Session.KeepAlive = true
		u.Session.StartHandle(c.Writer, c.Request)
	})

	e.Run("localhost:8080")
}
