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
	config.Add("app", config.StrMap{
		// 应用名称
		"name": config.Env("APP_NAME", "my_blog"),
		// APP URL
		"url": config.Env("APP_URL", "http://localhost:3000"),
		// 当前环境
		"env": config.Env("APP_ENV", "production"),
		// 是否进入调试模式
		"debug": config.Env("APP_DEBUG", false),
		// 应用服务端口
		"port": config.Env("APP_PORT", 3000),
		// SESSION 在 cookie 中加密的数据
		"key": config.Env("APP_KEY", "7a3b6f8b4b22f2bf13290104ec8b8a50787ad25f"),
	})
}
