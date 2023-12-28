// BaseController 基类控制器
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/28

package controllers

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
	"myblog/pkg/flash"
	"myblog/pkg/logger"
	"myblog/pkg/view"
)

// BaseController 基类控制器
type BaseController struct {
}

// ResponseForSQLError 处理 SQL 错误并返回
func (bc BaseController) ResponseForSQLError(w http.ResponseWriter, err error, notFoundMessage string, internalMessage string) {
	// 初始化消息提示
	if notFoundMessage == "" {
		notFoundMessage = "页面未找到"
	}

	if internalMessage == "" {
		internalMessage = "服务器内部错误"
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// SQL 未找到数据记录
		w.WriteHeader(http.StatusNotFound)
		view.RenderSimple(w, view.D{
			"Message": notFoundMessage,
		}, "errors.404")
	} else {
		// SQL 查询不到报错
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		view.RenderSimple(w, view.D{
			"Message": internalMessage,
		}, "errors.50x")
	}
}

// ResponseForUnauthorized 未授权的操作
func (bc BaseController) ResponseForUnauthorized(w http.ResponseWriter, r *http.Request) {
	flash.Warning("没有权限进行此操作")
	http.Redirect(w, r, "/", http.StatusFound)
}
