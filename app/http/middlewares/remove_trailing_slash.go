// middlewares
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package middlewares

import (
	"net/http"
	"strings"
)

// RemoveTrailingSlash 除首页外，移除所有请求路径后面的斜杠
func RemoveTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 除首页外
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}

		next.ServeHTTP(w, r)
	})
}
