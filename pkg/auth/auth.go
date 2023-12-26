// auth
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/23

package auth

import (
	"errors"

	"gorm.io/gorm"
	"myblog/app/models/user"
	"myblog/pkg/session"
)

func _getUID() string {
	_uid := session.Get("uid")
	uid, ok := _uid.(string)
	if ok && len(uid) > 0 {
		return uid
	}
	return ""
}

// User 获取登录用户信息
func User() user.User {
	uid := _getUID()
	if len(uid) > 0 {
		_user, err := user.Get(uid)
		if err == nil {
			return _user
		}

	}
	return user.User{}
}

// Attempt 尝试登录
func Attempt(email string, password string) error {
	// 根据 Email 获取用户
	_user, err := user.GetByEmail(email)
	// 如果出现 错误
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("账号不存在或者密码错误")
		} else {
			return errors.New("服务器内部错误，请稍后尝试")
		}
	} else {
		// 匹配密码
		if !_user.ComparePassword(password) {
			return errors.New("账号不存在或者密码错误")
		}
		// 登录用户，保存会话
		session.Put("uid", _user.GetStringID())
		return nil
	}
}

// Login 登录指定用户
func Login(_user user.User) {
	session.Put("uid", _user.GetStringID())
}

// Logout 退出指定用户
func Logout() {
	session.Forget("uid")
}

// Check 检测是否登录
func Check() bool {
	return len(_getUID()) > 0
}
