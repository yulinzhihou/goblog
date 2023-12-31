// Crud 文章增删改查
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package article

import (
	"net/http"

	"myblog/pkg/logger"
	"myblog/pkg/model"
	"myblog/pkg/pagination"
	"myblog/pkg/route"
	"myblog/pkg/types"
)

// Get 通过 ID 获取文章
func Get(idStr string) (Article, error) {
	var article Article
	id := types.StringToUint64(idStr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

// GetAll 获取文章列表
func GetAll(r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {
	// 初始化分页实例
	db := model.DB.Model(Article{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("articles.index"), perPage)

	// 获取实图数据
	viewData := _pager.Paging()
	// 获取数据
	var articles []Article
	err := _pager.Results(&articles)
	if err != nil {
		return nil, pagination.ViewData{}, err
	}

	return articles, viewData, nil
}

// Create 创建文章，通过 article.ID 来判断是否新增成功
func (article *Article) Create() (err error) {
	result := model.DB.Create(&article)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

// Update 更新文章
func (article *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}

// Delete 删除文章
func (article *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}

// GetByUserID 通过 user_id 获取用户所有的文章
func GetByUserID(uid string) ([]Article, error) {

	var articles []Article
	if err := model.DB.Where("user_id = ?", uid).Preload("User").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

// GetByCategoryID 通过 category_id 获取文章数据
func GetByCategoryID(categoryId string, r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {
	// 初始化分页数据
	db := model.DB.Model(Article{}).Where("category_id = ?", categoryId).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("categories.show", "id", categoryId), perPage)
	// 获取视图数据
	viewData := _pager.Paging()
	// 获取数据
	var articles []Article
	err := _pager.Results(&articles)
	if err != nil {
		return nil, pagination.ViewData{}, err
	}

	return articles, viewData, nil
}
