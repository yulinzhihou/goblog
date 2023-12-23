// requests
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/23

package requests

import (
	`github.com/thedevsaddam/govalidator`
	`myblog/app/models/user`
)

// ValidateRegistrationForm 验证表单， 返回 errs 长度等于零即验证通过
func ValidateRegistrationForm(data user.User) map[string][]string {
	// 定制誰规则
	rules := govalidator.MapData{
		"username":         []string{"required", "alpha_num", "between:3,30"},
		"email":            []string{"required", "min:4", "email"},
		"password":         []string{"required", "min:4", "max:30"},
		"password_confirm": []string{"required"},
	}

	// 定义错误信息
	messages := govalidator.MapData{
		"username": []string{
			"required:用户名必填项",
			"alpha_num:用户名格式错误，只允许数字和字母",
			"between:用户名的长度需介于 3-30 之间",
		},
		"email": []string{
			"required:邮箱必填项",
			"min:邮箱长度不能低于 4 位",
			"email:邮箱格式不对",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码最小长度为 4 位",
			"max:密码最大长度为 30 位",
		},
		"password_confirm": []string{
			"required:确认密码必填项",
		},
	}

	// 配置初始化
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	// 开始验证
	errs := govalidator.New(opts).ValidateStruct()
	// 因 govalidator 不支持 password_confirm 验证，自己单独写
	if data.Password != data.PasswordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入的密码不一样")
	}

	return errs
}
