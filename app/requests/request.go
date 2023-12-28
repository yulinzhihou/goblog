// requests
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/23

package requests

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

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

	// max_cn:2 最大中文长度
	govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		// 去除验证字段前缀，获取验证字段的值
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
		if valLength < l {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度需要小于 %d 个字", l)
		}
		return nil
	})

	// min_cn:2 最小中文字长度
	govalidator.AddCustomRule("min:cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		// 去除验证规则前缀，拿到验证字段的限制值
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min:cn"))
		// 进行判断
		if valLength < l {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度需要大于 %d 个字", l)
		}

		return nil
	})

}
