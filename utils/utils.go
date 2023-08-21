package utils

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

// 将map中的各个参数填充到url中
func FillURLWithMap(urlStr string, data map[string]interface{}) string {
	// 创建 URL 的参数
	params := url.Values{}
	for key, value := range data {
		params.Add(key, fmt.Sprintf("%v", value))
	}

	// 拼接参数到 URL
	return urlStr + "?" + params.Encode()
}

// 在两个整数中间随机随眠若干秒
func RandomSleep(minSeconds, maxSeconds int) {
	rand.Seed(time.Now().UnixNano())
	randomSeconds := float64(minSeconds+maxSeconds-minSeconds+1) + rand.Float64()
	duration := time.Duration(randomSeconds) * time.Second
	fmt.Printf("Sleeping for %f seconds...\n", randomSeconds)
	time.Sleep(duration)
	fmt.Println("Done sleeping!")
}
