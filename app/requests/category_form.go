// requests
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/28

package requests

import (
	"github.com/thedevsaddam/govalidator"
	"myblog/app/models/category"
)

// 定义规则
var categoryRules = govalidator.MapData{
	"name": []string{
		"required",
		"min_cn:2",
		"max_cn:20",
		"not_exists:categories,name"},
	"desc": []string{
		"max_cn:200",
	},
	"sort": []string{"between:0,999"},
}

// 定义错误消息
var categoryMessages = govalidator.MapData{
	"name": []string{
		"required:分类名称必填",
		"min_cn:分类名称至少要 2 个字",
		"max_cn:分类名称最多 20 个字",
		"not_exists:分类名称已经存在，请勿重复添加",
	},
	"desc": []string{
		"max_cn:分类描述最多 200 个字",
	},
	"sort": []string{
		"between:分类排序需介于 0-999 之间",
	},
}

func ValidateCategoryForm(data category.Category) map[string][]string {
	// 初始化配置
	opts := govalidator.Options{
		Data:          &data,
		Rules:         categoryRules,
		Messages:      categoryMessages,
		TagIdentifier: "valid",
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
