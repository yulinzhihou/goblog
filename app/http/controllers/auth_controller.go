// controllers
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/22

package controllers

import (
	"fmt"
	"net/http"

	"myblog/app/models/user"
	"myblog/pkg/view"
)

// AuthController 用户认证类，注册登录，退出
type AuthController struct {
}

// Register 用户注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 用户注册处理逻辑
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 初始化变量
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	// 表单验证，。
	_user := user.User{
		Username: username,
		Password: password,
		Email:    email,
	}

	_user.Create()

	if _user.ID > 0 {
		fmt.Fprint(w, "创建成功，ID "+_user.GetStringID())
	} else {
		view.RenderSimple(w, view.D{
			"Message": "创建用户失败，请联系管理员",
		}, "errors.50x")
	}

	// 表单验证不成功
}

// Login 用户登录
func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.login")
}

// DoLogin 登录的业务逻辑
func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	//
}

// Forget 忘记密码
func (*AuthController) Forget(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.forget")
}

// DoForget 忘记密码逻辑
func (*AuthController) DoForget(w http.ResponseWriter, r *http.Request) {
	//
}

// SendEmail 发送邮件
func (*AuthController) SendEmail(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.sendemail")
}

// DoSendEmail 发送邮件逻辑
func (*AuthController) DoSendEmail(w http.ResponseWriter, r *http.Request) {
	//
}

// Logout 退出系统逻辑
func (*AuthController) Logout(w http.ResponseWriter, r *http.Request) {

}
