// view
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/22

package view

import (
	"html/template"
	"io"
	"path/filepath"
	"strings"

	"myblog/pkg/logger"
	"myblog/pkg/route"
)

// D 是 map[string]interface{} 的简写
type D map[string]interface {
}

// Render 渲染通用视图
func Render(w io.Writer, data interface{}, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

// RenderSimple 渲染简单模板视图
func RenderSimple(w io.Writer, data interface{}, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

// RenderTemplate 渲染视图
func RenderTemplate(w io.Writer, name string, data interface{}, tplFiles ...string) {
	// 设置模板相对路径
	viewDir := "resources/views/"
	// 遍历传参文件列表 Slice, 设置正确路径，支持 dir.filename 语法糖，将 articles.show 更改为 articles/show
	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	// 所有布局模板文件 Slice
	layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)
	// 在 slice 里面新增我们的目标文件
	allFiles := append(layoutFiles, tplFiles...)
	// 解析所有模板文件
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"RouteName2URL": route.Name2URL,
	}).ParseFiles(allFiles...)

	logger.LogError(err)
	// 渲染模板
	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}
