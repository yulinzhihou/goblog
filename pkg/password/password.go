// password
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/26

package password

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"myblog/pkg/logger"
)

// Hash 使用 bcrypt 对密码进行加密
func Hash(password string) string {
	// GenerateFromPassword 的第二个参数是 Cost 值。建议大于 12，数值越大，耗费时间越长
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogError(err)

	return string(bytes)
}

// CheckHash 对比明文密码和数据库的哈希值
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	logger.LogError(err)
	fmt.Print(err == nil)
	return err == nil
}

// IsHashed 判断字符串是否哈希过的数据
func IsHashed(str string) bool {
	// bcrypt 加密后的长度等于 60
	return len(str) == 60
}
