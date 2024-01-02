// requests
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/27

package requests

import (
	"github.com/thedevsaddam/govalidator"
	"myblog/app/models/article"
)

// 定义规则
var articleRules = govalidator.MapData{
	"title":       []string{"required", "min_cn:3", "max_cn:60", "not_exists:articles,title"},
	"content":     []string{"required", "min_cn:10"},
	"category_id": []string{"required"},
	"user_id":     []string{"min:1"},
}

// 定义提示
var articleMessages = govalidator.MapData{
	"title": []string{
		"required:文章标题必填项",
		"min_cn:文章标题最小长度为 3 个字",
		"max_cn:文章标题最大长度为 60 个字",
	},
	"content": []string{
		"required:文章内容必填",
		"min_cn:文章内容最小 10 个字",
	},
	"category_id": []string{
		"required:文章分类必选择",
	},
	"user_id": []string{
		"min:用户未登录，请重新登录",
	},
}

// ValidateArticleForm 验证文章表单，返回 errs 长度等于零通过
func ValidateArticleForm(data article.Article) map[string][]string {
	opts := govalidator.Options{
		Data:          &data,
		Rules:         articleRules,
		Messages:      articleMessages,
		TagIdentifier: "valid",
	}
	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
