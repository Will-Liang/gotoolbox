package timeutils

import (
	"github.com/duke-git/lancet/v2/datetime"
	"strings"
	"time"
)

// 获取当前时间的时间戳，长度为13位(以毫秒为单位)
func GetCurrentUnix() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 得到时分
func GetHourMin() string {
	parts := strings.Split(datetime.GetNowTime(), ":")
	hour := parts[0]
	min := parts[1]
	return hour + min
}

// 获取当前时间格式化后的字符串
func GetCurrentFormattedTime() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	return formattedTime
}
