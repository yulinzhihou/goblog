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
	Username string `gorm:"type:varchar(128);default:null;unique;comment:用户名"`
	Email    string `gorm:"type:varchar(128);not null;unique;comment:邮箱"`
	Password string `gorm:"type:varchar(128);not null;comment:密码"`
	Gender   string `gorm:"type:tinyint(3);not null; default:0;comment:性别0=保密1=男2=女"`
	Desc     string `gorm:"type:json;not null;comment:用户信息"`
	Others   string `gorm:"type:text;not null;comment:其他"`
	Nickname string `gorm:"type:varchar(128);not null;default:'';comment:昵称"`
	Status   uint8  `gorm:"type:tinyint(3) unsigned;not null;default:1;comment:用户状态0=禁用1=正常"`
}
