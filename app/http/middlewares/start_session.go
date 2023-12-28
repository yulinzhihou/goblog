// middlewares
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/23

package middlewares

import (
	"net/http"

	"myblog/pkg/session"
)

// StartSession 开启 session 会话控制
func StartSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 启动会话
		session.StartSession(w, r)
		// 继续请求
		next.ServeHTTP(w, r)
	})
}
