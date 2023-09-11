package utils

import (
	crypto_rand "crypto/rand"
	"encoding/base64"
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

// 生成指定长度的随机字符串
func GenerateRandomString(length int) (string, error) {
	// 计算生成的字节数
	byteLength := (length * 6) / 8 // base64 编码后，每个字符占 6 位，8 位字节可以编码为 6 位字符

	// 生成随机字节
	randomBytes := make([]byte, byteLength)
	_, err := crypto_rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// 使用 base64 编码生成的随机字节
	randomString := base64.URLEncoding.EncodeToString(randomBytes)

	// 截取指定长度的字符串
	return randomString[:length], nil
}
