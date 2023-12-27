// middlewares
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/27

package middlewares

import (
	"net/http"
)

// HttpHandlerFunc 简写
type HttpHandlerFunc func(w http.ResponseWriter, r *http.Request)
