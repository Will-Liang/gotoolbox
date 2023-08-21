package maputils

// 判断传入的key是否存在于 map中
// map[string]类型
func ExistKey(m map[string]interface{}, key string) bool {
	_, exists := m[key]
	return exists
}
