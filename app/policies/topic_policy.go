// policies
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/28

package policies

import (
	"myblog/app/models/article"
	"myblog/pkg/auth"
)

// CanModifyArticle 是否允许文章被修改
func CanModifyArticle(_article article.Article) bool {
	return auth.User().ID == _article.UserID
}
