// bootstrap
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package bootstrap

import (
	"time"

	"gorm.io/gorm"
	"myblog/app/models/article"
	"myblog/app/models/user"
	"myblog/pkg/logger"
	"myblog/pkg/model"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	// 建立数据库连接池
	db := model.ConnectDB()

	// 命令行打印数据库请求连接信息
	sqlDB, _ := db.DB()
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(25)
	// 设置连接的过期时间
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// 创建和维护数据库
	migration(db)
}

// Migration 数据库自动迁移
func migration(db *gorm.DB) {
	// 自动迁移
	err := db.AutoMigrate(
		&user.User{},
		&article.Article{},
	)
	if err != nil {
		logger.LogError(err)
		return
	}
}
