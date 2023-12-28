// controllers
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/28

package controllers

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
	"myblog/app/models/article"
	"myblog/app/models/user"
	"myblog/pkg/logger"
	"myblog/pkg/route"
	"myblog/pkg/view"
)

// UserController 用户控制器类型
type UserController struct {
}

// Show 用户展示方法
func (*UserController) Show(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 获取对应的文章数据
	_user, err := user.Get(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			view.RenderSimple(w, view.D{
				"Message": "文章未找到",
			}, "errors.404")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			view.RenderSimple(w, view.D{
				"Message": "服务器内部错误",
			}, "errors.50x")
		}
	} else {
		// 读取成功，显示用户所有的文章
		articles, err1 := article.GetByUserID(_user.GetStringID())
		if err1 != nil {
			logger.LogError(err1)
			w.WriteHeader(http.StatusInternalServerError)
			view.RenderSimple(w, view.D{
				"Message": "数据错误，服务器内部错误",
			}, "errors.50x")

		} else {
			view.Render(w, view.D{
				"Articles": articles,
			}, "articles.index")
		}
	}
}
