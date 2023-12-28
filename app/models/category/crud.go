// category
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/28

package category

import (
	"myblog/pkg/logger"
	"myblog/pkg/model"
)

func (c *Category) Store() (err error) {
	if err = model.DB.Create(&c).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
