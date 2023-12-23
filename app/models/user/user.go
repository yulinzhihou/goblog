// user
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/22

package user

import (
	"myblog/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel
	Username string `gorm:"type:varchar(128);default:null;unique;comment:用户名" valid:"username"`
	Email    string `gorm:"type:varchar(128);not null;unique;comment:邮箱" valid:"email"`
	Password string `gorm:"type:varchar(128);not null;comment:密码" valid:"password"`
	Gender   string `gorm:"type:tinyint(3);not null; default:0;comment:性别0=保密1=男2=女" valid:"gender"`
	Desc     string `gorm:"type:json;default:null;comment:用户信息" valid:"desc"`
	Others   string `gorm:"type:text;default:null;comment:其他信息" valid:"others"`
	Nickname string `gorm:"type:varchar(128);not null;default:'';comment:昵称" valid:"nickname"`
	Status   uint8  `gorm:"type:tinyint(3) unsigned;not null;default:1;comment:用户状态0=禁用1=正常" valid:"status"`
	// gorm  -  设置，表示 GORM 在读写过程中忽略些字段
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}
