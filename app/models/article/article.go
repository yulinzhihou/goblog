// article
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package article

import (
	"strconv"

	"myblog/app/models"
	"myblog/pkg/route"
)

// Article 模型
type Article struct {
	models.BaseModel
	Title   string
	Image   string
	Brief   string
	Content string
	Status  uint8
	UserID  uint64
}

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.Name2URL("article.show", "id", strconv.FormatUint(article.ID, 10))
}
