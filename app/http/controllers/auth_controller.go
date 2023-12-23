// controllers
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/22

package controllers

import (
	"net/http"

	`myblog/app/models/user`
	`myblog/app/requests`
	"myblog/pkg/view"
)

// AuthController 用户认证类，注册登录，退出
type AuthController struct {
}

// 用户表单验证
type userForm struct {
	Username        string `valid:"name"`
	Password        string `valid:"password"`
	Email           string `valid:"email"`
	PasswordConfirm string `valid:"password_confirm"`
}

// Register 用户注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 用户注册处理逻辑
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 初始化变量
	_user := user.User{
		Username:        r.PostFormValue("username"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	// 开始验证
	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {
		// 表单验证不成功
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
	} else {
		err := _user.Create()

		if err != nil {
			if _user.ID > 0 {
				view.RenderSimple(w, view.D{
					"Errors": errs,
					"User":   _user,
				}, "auth.login")
			} else {
				view.RenderSimple(w, view.D{
					"Message": "创建用户失败，请联系管理员",
					"Errors":  errs,
					"User":    _user,
				}, "errors.50x")
			}
		} else {
			view.RenderSimple(w, view.D{
				"Errors": errs,
				"User":   _user,
			}, "auth.login")
		}
	}
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
