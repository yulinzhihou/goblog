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

	"myblog/pkg/auth"
	"myblog/pkg/logger"
	"myblog/pkg/route"
)

// D 是 map[string]interface{} 的简写
type D map[string]interface {
}

// Render 渲染通用视图
func Render(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

// RenderSimple 渲染简单模板视图
func RenderSimple(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

// RenderTemplate 渲染视图
func RenderTemplate(w io.Writer, name string, data D, tplFiles ...string) {
	// 通用模板数据
	data["isLogined"] = auth.Check()
	// 生成模板文件
	allFiles := getTemplateFiles(tplFiles...)
	// 解析所有模板文件
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"RouteName2URL": route.Name2URL,
	}).ParseFiles(allFiles...)

	logger.LogError(err)
	// 渲染模板
	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}

// getTemplateFiles 获取模板文章
func getTemplateFiles(tplFiles ...string) []string {
	// 设置模板相对路径
	viewDir := "resources/views/"
	// 遍历传参文件列表 Slice，设置正确路径，支持 dir.filename
	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	// 所有布局模板文件 Slice
	layoutFiles, err := filepath.Glob(viewDir + "layout/*.gohtml")
	logger.LogError(err)

	// 合并所有文件
	return append(layoutFiles, tplFiles...)
}
