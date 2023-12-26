// controllers
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package controllers

import (
	"net/http"

	"myblog/pkg/view"
)

// PagesController 处理静态页面
type PagesController struct {
}

// Home 首页
func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "index.index")
}

// About 关于页
func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "about.index")
}

// Help 帮助页
func (*PagesController) Help(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "help.index")
}

// NotFound 404页面
func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	view.Render(w, view.D{}, "error.404")
}

// InternalServer 50x页面
func (*PagesController) InternalServer(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "errors.50x")
}
