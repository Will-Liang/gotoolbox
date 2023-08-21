package strutils

import (
	"fmt"
	"github.com/Will-Liang/gotoolbox/logutils"
	"github.com/Will-Liang/gotoolbox/sat"
	"regexp"
	"strconv"
	"strings"
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

// 将字符串切片中的每一个字符串按照分隔符切分
// Deprecated: SplitStrWithDelimiters需要调用的，不推荐直接调用
func SplitAndCleanValues(inputValues []string, delimiter string) []string {

	cleanedValues := []string{}
	for _, inputValue := range inputValues {
		subValues := strings.Split(inputValue, delimiter)
		for _, subValue := range subValues {
			cleanedSubValue := strings.TrimSpace(subValue)
			if cleanedSubValue != "" {
				cleanedValues = append(cleanedValues, cleanedSubValue)
			}
		}
	}
	return cleanedValues
}

// 使用分隔符将字符串分成切片
func SplitStrWithDelimiters(value string, delimiters []string) []string {
	//示例：delimiters = []string{";", "；", "\n", "\r"}
	valuesToProcess := []string{value}
	for _, delimiter := range delimiters {
		valuesToProcess = SplitAndCleanValues(valuesToProcess, delimiter)
	}
	return valuesToProcess
}
