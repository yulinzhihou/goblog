// types
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package types

import (
	"strconv"

	"myblog/pkg/logger"
)

// Int64ToString 将 int64 转换为 string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// StringToUint64 将 string 转换成 uint64
func StringToUint64(idStr string) uint64 {
	i, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		logger.LogError(err)
	}
	return i
}

// StringToInt64 将 string 转换成 uint64
func StringToInt64(idStr string) int64 {
	i, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.LogError(err)
	}
	return i
}

// Uint64ToString 将 uint64 转换为 string
func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}
