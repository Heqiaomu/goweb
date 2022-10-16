// Package viperfile
/**
 * @Author: sunyang
 * @Email: sunyang@hyperchain.cn
 * @Date: 2022/10/12
 */
package viperfile

func ViperContainsKey(key string) bool {
	keys := GetViper().AllKeys()

	if len(keys) == 0 {
		return false
	}
	for _, k := range keys {
		if k == key {
			return true
		}
	}
	return false
}

func DefaultViperString(key, defaultValue string) string {
	if ViperContainsKey(key) {
		return GetViper().GetString(key)
	}
	return defaultValue
}

func DefaultViperInt(key string, defaultValue int) int {
	if ViperContainsKey(key) {
		return GetViper().GetInt(key)
	}
	return defaultValue
}

func DefaultViperBool(key string, defaultValue bool) bool {
	if ViperContainsKey(key) {
		return GetViper().GetBool(key)
	}
	return defaultValue
}
