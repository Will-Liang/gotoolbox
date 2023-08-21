package jsonutils

import (
	"encoding/json"
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/tidwall/gjson"
	"gotoolbox/fileutils"
	"gotoolbox/logutils"
	"gotoolbox/sliceutils"
	"os"
	"reflect"
)

// 根据key从json文件中读取value
func GetValueWithKeyFromFile(path string, key string) string {
	if !fileutil.IsExist(path) {
		fmt.Println(logutils.GetLog("ERROR", path+"不存在"))
		return ""
	}

	// 读取文件内容
	fileContent, err := fileutil.ReadFileToString(path)
	if err != nil {
		fmt.Println(logutils.GetLog("ERROR", err.Error()))
		return ""
	}

	// 判断文件是否为有效的JSON格式
	if !gjson.Valid(fileContent) {
		fmt.Println(logutils.GetLog("ERROR", path+" is not a valid JSON format."))
		return ""
	}
	result := gjson.Get(fileContent, key)
	return result.String()

}

// 将json写入到文件，标准json格式
func JsonToFile(path string, data interface{}) bool {
	fileutils.CheckPath(path, 1)
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(logutils.GetLog("ERROR", err.Error()))
		return false
	}

	err = fileutil.WriteBytesToFile(path, jsonData)
	if err != nil {
		fmt.Println(logutils.GetLog("ERROR", err.Error()))
		return false
	}

	return true
}

// 将json格式化后写入到文件，标准json格式
func JsonToFileFormat(path string, data interface{}) bool {
	fileutils.CheckPath(path, 1)
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(logutils.GetLog("ERROR", err.Error()))
		return false
	}

	err = fileutil.WriteBytesToFile(path, jsonData)
	if err != nil {
		fmt.Println(logutils.GetLog("ERROR", err.Error()))
		return false
	}

	return true
}

// 将json追加到文件中的最后一行
func JsonToLineFile(path string, data interface{}) bool {
	fileutils.CheckPath(path, 1)
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(logutils.GetLog("ERROR", err.Error()))
		return false
	}

	err = fileutil.WriteStringToFile(path, string(jsonData), true)
	if err != nil {
		fmt.Println(logutils.GetLog("ERROR", err.Error()))
		return false
	}

	return true
}

// 将结构体切片逐行写入到文件
// datas: [{a:"1"},{b:"2"}]
// 文件中:
// {a:"1"}
// {b:"2"}
func JsonSliceToLineFile(path string, datas interface{}) bool {
	//先判断传入的类型是不是切片
	if !sliceutils.IsSlice(datas) {
		fmt.Println(logutils.GetLog("ERROR", "datas is not a slice type"))
		return false
	}
	fileutils.CheckPath(path, 1)

	datas_slice := reflect.ValueOf(datas)
	for i := 0; i < datas_slice.Len(); i++ {
		json_data, err := json.Marshal(datas_slice.Index(i).Interface())
		if err != nil {
			fmt.Println(logutils.GetLog("ERROR", err.Error()))
			return false
		}
		err = fileutil.WriteStringToFile(path, string(json_data)+"\n", true)
		if err != nil {
			fmt.Println(logutils.GetLog("ERROR", err.Error()))
			return false
		}
	}

	return true
}

// 读取标准json格式，返回map
func FileToJsonMap(path string) (result map[string]interface{}) {
	//判断路径是否存在
	if !fileutil.IsExist(path) {
		fmt.Println(logutils.GetLog("ERROR", path+" not exist"))
		return nil
	}
	//读取文件是否有问题
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(logutils.GetLog("ERROR", err.Error()))
		return nil
	}

	if data == nil {
		return
	}

	//解析json是否有问题
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println(logutils.GetLog("ERROR", err.Error()))
		return nil
	}
	return

}

// 读取标准josn格式，返回
// object 传入: &object
func FileToJsonObject(path string, object interface{}) {
	fileutils.CheckPath(path, 1)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(logutils.GetLog("ERROR", err.Error()))
	}
	// 解析JSON文件并将其映射到给定的变量
	err = json.Unmarshal(data, object)
	if err != nil {
		fmt.Println(logutils.GetLog("ERROR", err.Error()))
	}
}
