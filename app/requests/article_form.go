// requests
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/27

package requests

import (
	"github.com/thedevsaddam/govalidator"
	"myblog/app/models/article"
)

// 定义规则
var articleRules = govalidator.MapData{
	"title": []string{"required"},
}

// 定义提示
var articleMessages = govalidator.MapData{}

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
