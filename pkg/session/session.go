// session
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/23

package session

import (
	"net/http"

	"github.com/gorilla/sessions"
	"myblog/pkg/logger"
)

// Store gorilla session 存储库
var Store = sessions.NewCookieStore([]byte("c58c936758f82aee259ba38591fb9cb6f23fa5df"))

// Session 当前会话
var Session *sessions.Session

// Request 用以获取用户会话
var Request *http.Request

// Response 用以写入用户会话
var Response http.ResponseWriter

// StartSession 初始化会话，在中间件中调用
func StartSession(w http.ResponseWriter, r *http.Request) {
	var err error
	// Store.Get() 的第二个参数 是 Cookie 的名称
	// gorilla/sessions 支持多会话，本项目我们只使用单一会话即可
	Session, err = Store.Get(r, "my-blog-session")
	logger.LogError(err)

	Request = r
	Response = w
}

// Put 写入键值对就的会话数据
func Put(key string, value interface{}) {
	Session.Values[key] = value
	Save()
}

// Get 获取会话数据，获取数据时请做类型检测
func Get(key string) interface{} {
	return Session.Values[key]
}

// Forget 删除某个会话项
func Forget(key string) {
	delete(Session.Values, key)
	Save()
}

// Flush 删除当前会话
func Flush() {
	Session.Options.MaxAge = -1
	Save()
}

// Save 保持会话
func Save() {
	// 非 HTTPS 的链接无法使用 Secure 和 httpOnly，浏览器会报错
	Session.Options.Secure = true
	Session.Options.HttpOnly = true
	Session.Options.SameSite = http.SameSiteNoneMode
	err := Session.Save(Request, Response)
	logger.LogError(err)
}