// config
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/27

package config

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"myblog/pkg/logger"
)

// Viper Viper 库实例
var Viper *viper.Viper

// StrMap 简写 -- map[string]interface{}
type StrMap map[string]interface{}

// init() 函数在 import 的时候立刻被加载
func init() {
	// 初始化 Viper 库
	Viper = viper.New()
	// 设置文件名称
	Viper.SetConfigName(".env")
	// 设置配置类型，支持 json .env yaml yml
	Viper.SetConfigType("env")
	// 配置环境亦是文件查找的路径，相对于 main.go
	Viper.AddConfigPath(".")
	// 开始读取目录下的 .env 文件，读取不到会报错
	err := Viper.ReadInConfig()
	logger.LogError(err)
	// 设置环境亦是前缀, 用以区分 go 的系统环境变量
	Viper.SetEnvPrefix("app_env")
	// Viper.Get() 时， 优先读取环境变量
	Viper.AutomaticEnv()
}

// Env 读取环境亦是，支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}
	return Get(envName)
}

// Add 新增配置项
func Add(name string, configuration map[string]interface{}) {
	Viper.Set(name, configuration)
}

// Get 获取配置项，允许使用点式获取，如：app.name
func Get(path string, defaultValue ...interface{}) interface{} {
	// 不存在
	if !Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}

	return Viper.Get(path)
}

// GetString 获取 string 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置数据
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

// GetInt64 获取 int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(path, defaultValue...))
}

// GetUint 获取 uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(Get(path, defaultValue...))
}

// GetUint64 获取 uint64 类型的配置信息
func GetUint64(path string, defaultValue ...interface{}) uint64 {
	return cast.ToUint64(Get(path, defaultValue...))
}

// GetBool 获取 bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}
