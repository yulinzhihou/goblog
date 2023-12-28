// category
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/28

package category

import (
	"myblog/pkg/logger"
	"myblog/pkg/model"
	"myblog/pkg/types"
)

// Store 分类新增保存
func (c *Category) Store() (err error) {
	if err = model.DB.Create(&c).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

// All 所有分类数据
func All() ([]Category, error) {
	var categories []Category
	if err := model.DB.Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

// Get 获取分类数据
func Get(idStr string) (Category, error) {
	var category Category
	id := types.StringToInt64(idStr)
	if err := model.DB.First(&category, id).Error; err != nil {
		return category, err
	}
	return category, nil
}
