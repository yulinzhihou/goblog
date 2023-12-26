// user
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/26

package user

import (
	"gorm.io/gorm"
	"myblog/pkg/password"
)

// BeforeCreate GORM 的模型勾子，创建模型应用前
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Password = password.Hash(user.Password)
	return
}

// BeforeUpdate GORM 的模型钩子，更新模型前调用
func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if !password.IsHashed(user.Password) {
		user.Password = password.Hash(user.Password)
	}
	return
}

// BeforeSave GORM 的模型钩子，在保存和更新模型前调用
func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if !password.IsHashed(user.Password) {
		user.Password = password.Hash(user.Password)
	}
	return
}
