// UserController 用户控制器
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/28

package controllers

import (
	"fmt"
	"net/http"

	"myblog/app/models/user"
	"myblog/pkg/route"
	"myblog/pkg/view"
)

// UserController 用户控制器类型
type UserController struct {
	BaseController
}

// Show 用户展示方法
func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 获取对应的文章数据
	_user, err := user.Get(id)
	fmt.Println(_user)
	fmt.Println(err)
	if err != nil {
		uc.ResponseForSQLError(w, err, "", "")
	} else {
		// 读取成功，显示用户所有的文章
		// articles, err1 := article.GetByUserID(_user.GetStringID())
		// if err1 != nil {
		// 	uc.ResponseForSQLError(w, err, "", "")
		// } else {
		view.Render(w, view.D{
			"User": _user,
		}, "users.show")
		// }
	}
}

// Update 用户基础资料更新，不包括密码
func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	//
}
