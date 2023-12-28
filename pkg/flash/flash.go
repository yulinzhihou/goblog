// flash
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/27

package flash

import (
	"encoding/gob"

	"myblog/pkg/session"
)

// Flashes Flash 消息数据类型，用以在会话中存在 map
type Flashes map[string]interface{}

// 存入会话数据里面的 Key
var flashKey = "_flashes"

func init() {
	// 在 gorilla/sessions 中存储 map 和 struct 数据需要提前注册 gob, 方便后续 gob 序列化编码，解码
	gob.Register(Flashes{})
}

// Info 添加 info 类型的消息提示
func Info(message string) {
	addFlashes("info", message)
}

// Warning 添加 warning 类型的消息提示
func Warning(message string) {
	addFlashes("warning", message)
}

// Success 添加 success 类型的消息提示
func Success(message string) {
	addFlashes("success", message)
}

// Danger 添加 danger 类型的消息提示
func Danger(message string) {
	addFlashes("danger", message)
}

// All 获取所有消息
func All() Flashes {
	val := session.Get(flashKey)
	// 读取时必须做类型验证
	flashMessages, ok := val.(Flashes)

	if !ok {
		return nil
	}
	// 读取即销毁，直接删除
	session.Forget(flashKey)
	return flashMessages
}

// 私有方法，新增一条提示消息
func addFlashes(key string, message string) {
	flashes := Flashes{}
	flashes[key] = message
	session.Put(flashKey, flashes)
	session.Save()
}
