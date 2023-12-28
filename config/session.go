// config
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/27

package config

import (
	"myblog/pkg/config"
)

func init() {
	config.Add("session", config.StrMap{
		// 目前只支持 Cookie
		"default": config.Env("SESSION_DRIVER", "cookie"),
		// 会话 Cookie 名称
		"session_name": config.Env("SESSION_NAME", "my_blog_session"),
	})
}
