// category
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/28

package category

import (
	"myblog/app/models"
)

// Category 文章分类 GORM
type Category struct {
	models.BaseModel
	Name string `gorm:"type:varchar(255);not null;comment:分类名称" valid:"name"`
	Desc string `gorm:"type:varchar(255);default:null;comment:分类描述" valid:"desc"`
	Sort int    `gorm:"type:int(11) unsigned;not null;default:0;comment:分类排序" valid:"sort"`
}
