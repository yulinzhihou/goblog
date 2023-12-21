// database
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package database

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"myblog/pkg/logger"
)

// DB 数据库对象
var DB *sql.DB

// Initialize 初始化数据库
func Initialize() {
	initDB()
	createTables()
}

func initDB() {
	var err error

	// 设置数据库连接信息
	config := mysql.Config{
		User:                 "root",
		Passwd:               "orico@f2b211.com",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "myblog",
		AllowNativePasswords: true,
	}
	// 准备数据库连接池
	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	// 设置最大连接数
	DB.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	DB.SetMaxIdleConns(25)
	// 设置每个连接的超时时间
	DB.SetConnMaxLifetime(5 * time.Minute)
	// 尝试连接，失败报错
	err = DB.Ping()
	logger.LogError(err)
}

func createTables() {
	createArticlesSQL := `
	create table if not exists articles (
	    id bigint(20) primary key auto_increment not null comment "ID",
	    user_id bigint(20) unsigned not null comment "用户ID",
	    image varchar(255) collate utf8mb4_unicode_ci not null comment "图片",
	    bref varchar(255) collate utf8mb4_unicode_ci not null comment "文章简介",
	    title varchar(255) collate utf8mb4_unicode_ci not null comment "文章标题",
	    content longtext collate utf8mb4_unicode_ci not null comment "文章内容",
	    status tinyint(3) unsigned not null default 0 comment "文章状态0=草稿1=已发布2=垃圾桶"
	) engine="innodb" default charset utf8mb4 collate utf8mb4_unicode_ci comment "文章表";
`
	_, err := DB.Exec(createArticlesSQL)
	logger.LogError(err)
}
