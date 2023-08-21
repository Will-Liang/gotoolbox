package sliceutils

import "reflect"

// 判断传入的类型是不是切片
func IsSlice(data interface{}) bool {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		return false
	}
	return true
}
