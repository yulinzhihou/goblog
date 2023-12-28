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
	config.Add("database", config.StrMap{
		// 数据库连接池信息
		"mysql": map[string]interface{}{
			"host":     config.Env("DB_HOST", "127.0.0.1"),
			"port":     config.Env("DB_PORT", 3306),
			"database": config.Env("DB_DATABASE", "myblog"),
			"username": config.Env("DB_USER", "root"),
			"password": config.Env("DB_PASS", "123456"),
			"charset":  config.Env("DB_CHARSET", "utf8mb4"),
			// 连接池配置
			"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 25),
			"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 100),
			"max_life_seconds":     config.Env("DB_MAX_LIFE_SECONDS", 5*60),
		},
		// redis 配置
		"redis": map[string]interface{}{
			"host":    config.Env("REDIS_HOST", "127.0.0.1"),
			"port":    config.Env("REDIS_PORT", 6379),
			"pass":    config.Env("REDIS_PASS", ""),
			"select":  config.Env("REDIS_SELECT", 0),
			"timeout": config.Env("REDIS_TIMEOUT", 120),
			"prefix":  config.Env("REDIS_PREFIX", ""),
		},
	})
}
