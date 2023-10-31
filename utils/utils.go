package utils

import (
	"crypto/md5"
	crypto_rand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/Will-Liang/gotoolbox/logutils"
	"math/rand"
	"net/url"
	"reflect"
	"strings"
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

// 将结构体中的各个参数填充到url中
// flag: 0 默认，1 表示将字段名转换为小写
// 如果是指针，应该写为 *s
func FillURLWithStruct(urlStr string, s interface{}, flag int) string {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		logutils.PrintErrorLog("The parameter passed in is not a structure. The parameter.Kind() is" + v.Kind().String())
		return "" // 返回空字符串表示不是结构体
	}

	u := url.Values{}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := v.Type().Field(i).Name
		fieldValue := fmt.Sprintf("%v", field.Interface())

		if fieldValue != "" && flag == 1 {
			// 将字段名转为小写
			lowercaseFieldName := strings.ToLower(fieldName)
			u.Add(lowercaseFieldName, fieldValue)
		}
		if fieldValue != "" && flag == 0 {
			u.Add(fieldName, fieldValue)
		}

	}

	return urlStr + "?" + u.Encode()
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
// length 长度
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

// 计算字符串的 MD5 哈希值并返回十六进制字符串
func GenerateStrMD5(input string) string {
	// 创建 MD5 哈希对象
	hasher := md5.New()

	// 将字符串转换为字节数组并写入哈希对象
	hasher.Write([]byte(input))

	// 计算哈希值并将其转换为十六进制字符串
	md5Hash := hasher.Sum(nil)
	md5String := hex.EncodeToString(md5Hash)
	// 长度 32位
	return md5String
}
