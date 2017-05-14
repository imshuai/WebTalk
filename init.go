package main

import (
	"os"

	"github.com/imshuai/lightlog"
)



func init() {
	DatabaseInit()
	var delimiter string
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		delimiter = "\\"
	} else {
		delimiter = "/"
	}
	dir, _ := os.Getwd() //当前的目录

	lightlog.SetLevel(lightlog.INFO)
	lightlog.SetPrefix("[WebTalk]")
	lightlog.SetRollingFile(dir+delimiter+"logs", "server.log", 10, 1, lightlog.MB)
}
