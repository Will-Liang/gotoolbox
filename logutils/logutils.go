package logutils

import (
	"fmt"
	"runtime"
	"time"
)

// 获得一个日志信息
// message 信息
func GetLog(log_type string, message string) string {
	_, file, line, _ := runtime.Caller(1)
	currentTime := time.Now()
	standardTime := currentTime.Format("2006-01-02 15:04:05")
	s := fmt.Sprintf("\n=======================================\n"+
		"%s (%s) Location: %s:%d\n%s\n", log_type, standardTime, file, line, message)
	return s
}
