package main

import (
	"net/http"

	"myblog/app/http/middlewares"
	"myblog/bootstrap"
	"myblog/config"
	c "myblog/pkg/config"
	"myblog/pkg/logger"
)

// 初始化配置文件信息
func init() {
	config.Initialize()
}

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
