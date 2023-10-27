package logutils

import (
	"fmt"
	"runtime"
	"strings"
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

// 获得带有颜色的日志信息
// 如果直接调用该方法，建议将 skip 设置为 1
func GetColorLog(log_type string, message string, skip int) string {
	//目前仅支持log_type 为 ERROR,WARNING,INFO
	_, file, line, _ := runtime.Caller(skip)
	currentTime := time.Now()
	standardTime := currentTime.Format("2006-01-02 15:04:05")
	log_type = strings.ToUpper(log_type)
	s := ""
	if log_type == "ERROR" {
		s = fmt.Sprintf("\x1b[31m"+
			"%s (%s) Location: %s:%d\n%s\n\x1b[0m", log_type, standardTime, file, line, message)
	} else if log_type == "WARNING" {
		s = fmt.Sprintf("\x1b[33m"+
			"%s (%s) Location: %s:%d\n%s\n\x1b[0m", log_type, standardTime, file, line, message)
	} else if log_type == "INFO" {
		s = fmt.Sprintf("\x1b[32m"+
			"%s (%s) Location: %s:%d\n%s\n\x1b[0m", log_type, standardTime, file, line, message)
	} else {
		s = fmt.Sprintf("%s (%s) Location: %s:%d\n%s\n", log_type, standardTime, file, line, message)
	}
	return s

}

// 打印错误日志信息
func PrintErrorLog(message string) {
	printColorLog("ERROR", message)
}

// 打印警告日志信息
func PrintWarningLog(message string) {
	printColorLog("WARNING", message)
}

// 打印提示日志信息
func PrintInfoLog(message string) {
	printColorLog("INFO", message)
}

// 打印带有颜色的日志信息
func printColorLog(log_type string, message string) {
	fmt.Println(GetColorLog(log_type, message, 3))
}
