// logger
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package logger

import (
	"log"
)

// LogError 当存在错误记录日志
func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}
