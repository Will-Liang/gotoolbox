package strutils

import (
	"fmt"
	"github.com/Will-Liang/gotoolbox/logutils"
	"github.com/Will-Liang/gotoolbox/sat"
	"regexp"
	"strconv"
)

// 传入字符串，判断字符串是否纯英文
func IsAlphabet(str string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z]+$", str)
	return match
}

// 繁体转简体
func TraditionalConvertSimple(str string) (result string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(logutils.GetLog("ERROR", "繁体转简体发生错误"))
		}
	}()
	result = str
	dicter := sat.DefaultDict()
	result = dicter.Read(str)

	return result
}

// 简体转繁体
func SimpleConvertTraditional(str string) (result string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(logutils.GetLog("ERROR", "简体转繁体发生错误"))

		}
	}()
	result = str
	dicter := sat.DefaultDict()
	result = dicter.ReadReverse(str)

	return result
}

// 将字符串转成float
func StrToFloat(str string) float32 {
	value, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0.0
	}
	return float32(value)
}

// 将字符串转成Int
func StrToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return value
}
