// Article 文章模型
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package article

import (
	"myblog/app/models"
	"myblog/app/models/category"
	"myblog/app/models/user"
	"myblog/pkg/route"
	"myblog/pkg/types"
)

// Article 模型
type Article struct {
	models.BaseModel
	Title      string            `gorm:"type:varchar(255);not null;default:'';comment:标题" valid:"title"`
	Image      string            `gorm:"type:varchar(255);not null;default:'';comment:图片" valid:"image"`
	Brief      string            `gorm:"type:varchar(255);not null;default:'';comment:简介" valid:"brief"`
	Content    string            `gorm:"type:longtext;comment:内容" valid:"content"`
	Status     uint8             `gorm:"type:tinyint unsigned;true;not null;default:0;comment:状态0=草稿1=已发布2=未发布" valid:"status"`
	UserID     uint64            `gorm:"type:bigint unsigned;not null;index;comment:用户ID" valid:"user_id"`
	User       user.User         `gorm:"-"`
	CategoryID uint64            `gorm:"type:bigint unsigned; not null;default:0; index; comment:分类ID" valid:"category_id"`
	Category   category.Category `gorm:"-"`
}

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.Name2URL("article.show", "id", types.Uint64ToString(article.ID))
}

// CreatedAtDate 创建日期格式化
func (article Article) CreatedAtDate() string {
	return article.CreatedAt.Format("2023-12-21")
}
