// controllers
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/28

package controllers

import (
	"net/http"

	"github.com/spf13/cast"
	"myblog/app/models/article"
	"myblog/app/models/category"
	"myblog/app/requests"
	"myblog/pkg/flash"
	"myblog/pkg/route"
	"myblog/pkg/view"
)

// CategoriesController 文章分类控制器
type CategoriesController struct {
	BaseController
}

// Index 文章分类列表
func (cc CategoriesController) Index(w http.ResponseWriter, r *http.Request) {

}

// Show 文章分类详情
func (cc CategoriesController) Show(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 读取对应的分类数据
	_category, err := category.Get(id)
	// 获取结果集
	articles, pagerData, err := article.GetByCategoryID(_category.GetStringID(), r, 2)
	// 判断
	if err != nil {
		cc.ResponseForSQLError(w, err, "", "")
	} else {
		view.Render(w, view.D{
			"Articles":  articles,
			"PagerData": pagerData,
		}, "articles.index")
	}
}

// Create 文章分类创建
func (cc CategoriesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "categories.create", "categories._form_field", "categories._form_slide")
}

// Store 文章分类保存
func (cc *CategoriesController) Store(w http.ResponseWriter, r *http.Request) {
	// 初始化数据
	_category := category.Category{
		Name: r.PostFormValue("name"),
		Desc: r.PostFormValue("desc"),
		Sort: cast.ToInt(r.PostFormValue("sort")),
	}
	// 表单验证
	errs := requests.ValidateCategoryForm(_category)
	// 开始判断是否通过验证
	if len(errs) == 0 {
		// 创建文章分类
		err := _category.Store()
		// 根据返回的 ID
		if err != nil && _category.ID > 0 {
			flash.Success("创建成功")
			indexURL := route.Name2URL("home")
			http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			cc.ResponseForSQLError(w, err, "", "创建文章分类失败，请联系管理员")
		}
	} else {
		view.Render(w, view.D{
			"Category": _category,
			"Errors":   errs,
		}, "categories.create", "categories._form_field", "categories._form_slide")
	}
}
