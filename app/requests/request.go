// requests
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/23

package requests

import (
	"errors"
	"fmt"
	"strings"

	"myblog/pkg/model"

	"github.com/thedevsaddam/govalidator"
)

// 此方法会在初始化时执行
func init() {
	// not_exists:users,email
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")
		tableName := rng[0]
		dbField := rng[1]
		val := value.(string)

		var count int64
		model.DB.Table(tableName).Where(dbField+" = ?", val).Count(&count)
		if count != 0 {
			if message != "" {
				return errors.New(message)
			}

			return fmt.Errorf("%v 已被占用", val)
		}

		return nil

	})
}
