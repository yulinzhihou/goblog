// config
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/28

package config

import (
	"myblog/pkg/config"
)

func init() {
	config.Add("pagination", config.StrMap{
		// 默认每页条数
		"per_page": 10,
		// URL 中用来分辨多少页的参数
		"url_query": "page",
	})
}
