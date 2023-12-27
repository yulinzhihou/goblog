// controllers
// Author : yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/22

package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
	"myblog/app/models/user"
	"myblog/app/requests"
	"myblog/pkg/auth"
	"myblog/pkg/flash"
	"myblog/pkg/password"
	"myblog/pkg/route"
	"myblog/pkg/view"
)

// AuthController 用户认证类，注册登录，退出
type AuthController struct {
}

// Register 用户注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {

	view.RenderSimple(w, view.D{
		"User": user.User{},
	}, "auth.register")
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
				// 登录用户并跳转到首页，提示消息
				flash.Success("恭喜注册成功！")
				auth.Login(_user)
				http.Redirect(w, r, route.Name2URL("auth.login"), http.StatusFound)
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
	view.RenderSimple(w, view.D{
		"User":   user.User{},
		"Errors": nil,
	}, "auth.login")
}

// DoLogin 登录的业务逻辑
func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	// 初始化表单
	_email := r.PostFormValue("email")
	_password := r.PostFormValue("password")

	if err := auth.Attempt(_email, _password); err == nil {
		// 登录成功。
		flash.Success("欢迎回来！")
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		view.RenderSimple(w, view.D{
			"Error":    err,
			"Email":    _email,
			"Password": _password,
		}, "auth.login")
	}
}

// Forget 忘记密码
func (*AuthController) Forget(w http.ResponseWriter, r *http.Request) {
	// 获取重置密码链接
	id := r.FormValue("id")
	// 查询用户数据
	_user, err := user.Get(id)
	// 判断用户是否存在
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			view.RenderSimple(w, view.D{
				"Message": "用户数据不存在",
			}, "errors.404")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			view.RenderSimple(w, view.D{
				"Message": "服务器内部错误",
			}, "errors.50x")
		}
	} else {
		view.RenderSimple(w, view.D{
			"User": _user,
		}, "auth.forget")
	}

}

// ResetPassword 忘记密码逻辑
func (*AuthController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	// 验证邮件链接，解析出ID和唯一码，现在没有这个唯一码，暂时先实现逻辑
	id := r.PostFormValue("id")
	_user, err := user.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			view.RenderSimple(w, view.D{
				"Message": "用户未找到",
			}, "errors.404")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			view.RenderSimple(w, view.D{
				"Message": "服务器内部错误",
			}, "errors.50x")
		}
	} else {
		// 表示查询用户成功。
		_user1 := user.User{
			Password:        r.PostFormValue("password"),
			PasswordConfirm: r.PostFormValue("password_confirm"),
		}
		errs := requests.ValidateResetPasswordFormData(_user1)
		if len(errs) == 0 {
			view.RenderSimple(w, view.D{
				"Errors": errs,
				"User":   _user,
			}, "auth.forget")
		} else {
			_user.Password = password.Hash(_user1.Password)
			rowsAffected, err1 := _user.Update()

			if err1 != nil {
				// 数据库异常
				w.WriteHeader(http.StatusInternalServerError)
				view.RenderSimple(w, view.D{
					"Message": "服务器内部错误",
				}, "errors.50x")
				return
			}

			if rowsAffected > 0 {
				showURL := route.Name2URL("auth.login")
				http.Redirect(w, r, showURL, http.StatusFound)
			} else {
				_, err = fmt.Fprint(w, "您未作出任何修改")
				if err != nil {
					return
				}
			}
		}
	}

}

// SendEmail 发送邮件
func (*AuthController) SendEmail(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.send_email")
}

// DoSendEmail 发送邮件逻辑
func (*AuthController) DoSendEmail(w http.ResponseWriter, r *http.Request) {
	// 接收用户提交的 Email 数据
	email := r.PostFormValue("email")
	// 根据 Email 查询用户
	_user, err := user.GetByEmail(email)
	// 如果出现 错误
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			view.RenderSimple(w, view.D{
				"Messages": "账号不存在",
			}, "errors.404")
		} else {
			view.RenderSimple(w, view.D{
				"Message": "内部错误，请稍后再试",
			}, "errors.50x")
		}
	} else {
		// 表示没有问题
		id := _user.ID
		// TODO: 发送带 ID 链接邮件到邮箱
		w.WriteHeader(http.StatusOK)
		_, err2 := fmt.Fprint(w, "邮件发送成功， ID :"+strconv.FormatUint(id, 10))
		if err2 != nil {
			return
		}
	}
}

// Logout 退出系统逻辑
func (*AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	auth.Logout()
	flash.Success("您已成功退出登录！")
	http.Redirect(w, r, "/", http.StatusFound)
}
