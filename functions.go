package main

import (
	"crypto/rand"
	"fmt"
	"io"

	"github.com/imshuai/lightlog"
)

//LogDebug 记录Debug级别日志
func LogDebug(v ...interface{}) {
	lightlog.Debug(v)
}

//LogInfo 记录Info级别日志
func LogInfo(v ...interface{}) {
	lightlog.Info(v)
}

//LogWarn 记录Warn级别日志
func LogWarn(v ...interface{}) {
	lightlog.Warn(v)
}

//LogError 记录Error级别日志
func LogError(v ...interface{}) {
	lightlog.Error(v)
}

func uuid() string {
	b := make([]byte, 16)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
