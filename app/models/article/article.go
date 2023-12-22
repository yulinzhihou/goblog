// article
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package article

import (
	"strconv"

	"myblog/app/models"
	"myblog/app/models/user"
	"myblog/pkg/route"
)

// Article 模型
type Article struct {
	models.BaseModel
	Title   string    `gorm:"type:varchar(255);not null;default:'';comment:标题"`
	Image   string    `gorm:"type:varchar(255);not null;default:'';comment:图片"`
	Brief   string    `gorm:"type:varchar(255);not null;default:'';comment:简介"`
	Content string    `gorm:"type:longtext;comment:内容"`
	Status  uint8     `gorm:"type:tinyint unsigned;true;not null;default:0;comment:状态0=草稿1=已发布2=未发布"`
	UserID  uint64    `gorm:"type:bigint unsigned;not null;unsigned;comment:用户ID"`
	User    user.User `gorm:"-"`
}

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.Name2URL("article.show", "id", strconv.FormatUint(article.ID, 10))
}
