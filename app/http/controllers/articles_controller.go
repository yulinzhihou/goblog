// controllers
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package controllers

import (
	"net/http"
	"strconv"

	"myblog/app/models/article"
	"myblog/app/policies"
	"myblog/app/requests"
	"myblog/pkg/flash"
	"myblog/pkg/route"
	"myblog/pkg/view"
)

// ArticlesController 文章相关页面
type ArticlesController struct {
	BaseController
}

// Index 文章列表页
func (ac *ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	// 获取结果集
	_articles, err := article.GetAll()

	if err != nil {
		ac.ResponseForSQLError(w, err, "", "")
	} else {
		view.Render(w, view.D{
			"Articles": _articles,
		}, "articles.index")
	}
}

// Create 文章创建表单页面
func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{
		"Articles": nil,
	}, "articles.create", "articles._form_slide", "articles._form_field")
}

// Store 文章新增
func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	_article := article.Article{
		Title:   r.PostFormValue("title"),
		Content: r.PostFormValue("content"),
	}

	// 验证器开始验证
	errs := requests.ValidateArticleForm(_article)

	// 检查是否有错误
	if len(errs) == 0 {
		err := _article.Create()
		if err != nil && _article.ID > 0 {
			showURL := route.Name2URL("article.index")
			flash.Success("插入成功，ID 为" + strconv.FormatUint(_article.ID, 10))
			http.Redirect(w, r, showURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			view.RenderSimple(w, view.D{
				"Message": "创建文章失败，服务器内部错误",
			}, "errors.50x")
		}
	} else {
		view.Render(w, view.D{
			"Article": _article,
			"Errors":  errs,
		}, "articles.create", "articles._form_field", "articles._form_slide")
	}
}

// Show 文章详情页面
func (ac *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 获取文章数据
	_article, err := article.Get(id)
	// 如果出现错误
	if err != nil {
		ac.ResponseForSQLError(w, err, "", "")
	} else {
		// 读取成功。
		view.Render(w, view.D{
			"Article":          _article,
			"CanModifyArticle": policies.CanModifyArticle(_article),
		}, "articles.show")
	}
}

// Edit 文章编辑页面
func (ac *ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 获取对应文章数据
	_article, err := article.Get(id)

	// 如果出现错误
	if err != nil {
		ac.ResponseForSQLError(w, err, "", "")
	} else {
		if !policies.CanModifyArticle(_article) {
			ac.ResponseForUnauthorized(w, r)
		} else {
			view.Render(w, view.D{
				"Article": _article,
			}, "articles.edit", "articles._form_slide", "articles._form_field")
		}
	}
}

// Update 文章修改
func (ac *ArticlesController) Update(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 获取对应文章数据
	_article, err := article.Get(id)
	// 判断是否有错误
	if err != nil {
		ac.ResponseForSQLError(w, err, "", "")
	} else {
		if !policies.CanModifyArticle(_article) {
			ac.ResponseForUnauthorized(w, r)
		} else {
			// 未出现错误
			_article1 := article.Article{
				Title:   r.PostFormValue("title"),
				Content: r.PostFormValue("content"),
			}

			errs := requests.ValidateArticleForm(_article1)

			if len(errs) == 0 {

				rowsAffected, err1 := _article.Update()

				if err1 != nil {
					// 数据库错误
					w.WriteHeader(http.StatusInternalServerError)
					view.RenderSimple(w, view.D{
						"Message": "服务器内部错误",
					}, "errors.50x")
					return
				}

				// 表示更新成功
				if rowsAffected > 0 {
					showURL := route.Name2URL("articles.show", "id", id)
					http.Redirect(w, r, showURL, http.StatusFound)
				}
			} else {
				// 表示表单验证不通过 ，显示错误信息
				view.Render(w, view.D{
					"Article": _article,
					"Errors":  errs,
				}, "articles.edit")
			}
		}
	}
}

// Delete 文章删除
func (ac *ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 读取对应的文章数据
	_article, err := article.Get(id)

	if err != nil {
		ac.ResponseForSQLError(w, err, "", "")
	} else {
		if !policies.CanModifyArticle(_article) {
			ac.ResponseForUnauthorized(w, r)
		} else {
			// 未出现错误，进行删除操作
			rowsAffected, err1 := _article.Delete()
			// 是否发生错误
			if err1 != nil {
				w.WriteHeader(http.StatusInternalServerError)
				view.RenderSimple(w, view.D{
					"Message": "服务器内部错误",
				}, "errors.50x")
			} else {
				// 未发生错误
				if rowsAffected > 0 {
					// 删除成功，跳转
					indexURL := route.Name2URL("articles.index")
					http.Redirect(w, r, indexURL, http.StatusFound)
				} else {
					w.WriteHeader(http.StatusNotFound)
					view.RenderSimple(w, view.D{
						"Message": "文章未找到",
					}, "errors.404")
				}
			}
		}

	}
}
