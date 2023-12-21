// middlewares
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package middlewares

import (
	"net/http"
)

// ForceHTML 强制标头返回 HTML 内容类型
func ForceHTML(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置响应头
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		// 继续请求
		next.ServeHTTP(w, r)
	})
}
