package main

import (
	"net/http"

	"myblog/app/http/middlewares"
	"myblog/bootstrap"
	"myblog/pkg/logger"
)

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
