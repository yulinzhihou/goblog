// article
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package article

// Article 模型
type Article struct {
	ID      uint64
	Title   string
	Image   string
	Brief   string
	Content string
	UserID  uint64
}
