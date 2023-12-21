package main

import (
	"net/http"

	"myblog/app/http/middlewares"
	"myblog/bootstrap"
	"myblog/pkg/logger"
)

// ArticlesFormData 创建博文表单数据
// type ArticlesFormData struct {
// 	Title, Body string
// 	URL         *url.URL
// 	Errors      map[string]string
// }

// func defaultHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html;charset=utf-8")
// 	if r.URL.Path == "/" {
// 		_, err := fmt.Fprint(w, "Hello goblog")
// 		if err != nil {
// 			return
// 		}
// 	} else {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "404 页面未找到")
// 	}
// }
//
// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	fmt.Fprint(w, "about")
// }
//
// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "首页")
// }

// helpHandler 帮助页
// func helpHandler(w http.ResponseWriter, r *http.Request) {
//
// }

// notFoundHandler 404 页面
// func notFoundHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	_, err := fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或者建议，请联系"+"<a href=\"mailto:yulinzhihou@163.com\">yulinzhihou@163.com</a>")
// 	if err != nil {
// 		return
// 	}
// }

/*
入口方法
*/
func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
