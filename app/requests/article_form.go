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
	"title":   []string{"required", "min:3", "max:60"},
	"content": []string{"required", "min:10"},
}

// 定义提示
var articleMessages = govalidator.MapData{
	"title": []string{
		"required:文章标题必填项",
		"min:文章标题最小长度为 3",
		"max:文章标题最大长度为 60",
	},
	"content": []string{
		"required:文章内容必填",
		"min:文章内容最小 10 个字符",
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
