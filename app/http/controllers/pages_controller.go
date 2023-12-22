// controllers
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package controllers

import (
	"fmt"
	"net/http"
)

// PagesController 处理静态页面
type PagesController struct {
}

// Home 首页
func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "首页首页首页")
}

// About 关于页
func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "关于页")
}

// Help 帮助页
func (*PagesController) Help(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "帮助页")
}

// NotFound 404页面
func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404")
}

// InternalServer 50x页面
func (*PagesController) InternalServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "50x 页面")
}
