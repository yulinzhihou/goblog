// controllers
// Author : Yulinzhihou
// Github : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"unicode/utf8"

	"gorm.io/gorm"
	"myblog/app/models/article"
	"myblog/pkg/logger"
	"myblog/pkg/route"
	"myblog/pkg/view"
)

// ArticlesController 文章相关页面
type ArticlesController struct {
}

type ArticlesFormData struct {
	ID                                uint64
	Image, Brief, Content, Title, URL string
	Status                            int8
	UserID                            uint64
	Errors                            map[string]string
}

// Index 文章列表页
func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	// 获取结果集
	_articles, err := article.GetAll()

	if err != nil {
		// 数据库错误
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	} else {
		view.Render(w, view.D{
			"Articles": _articles,
		}, "articles.index")
	}
}

// Create 文章创建表单页面
func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	// storeURL := route.Name2URL("articles.store")
	//
	// data := ArticlesFormData{
	// 	Image:  "",
	// 	URL:    storeURL,
	// 	Errors: nil,
	// }
	view.Render(w, view.D{
		"Articles": nil,
	}, "articles.create", "articles._form_slide", "articles._form_field")
}

// Store 文章新增
func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	content := r.PostFormValue("content")

	errors := validateArticleFormData(title, content)

	// 检查是否有错误
	if len(errors) == 0 {
		_article := article.Article{
			Title:   title,
			Content: content,
		}
		_article.Create()
		if _article.ID > 0 {
			fmt.Fprint(w, "插入成功，ID 为"+strconv.FormatUint(_article.ID, 10))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		storeURL := route.Name2URL("articles.store")

		data := ArticlesFormData{
			Title:   title,
			Content: content,
			URL:     storeURL,
			Errors:  errors,
		}

		tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
		logger.LogError(err)

		err = tmpl.Execute(w, data)
		logger.LogError(err)
	}
}

// Show 文章详情页面
func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 获取文章数据
	_article, err := article.Get(id)
	// 如果出现错误
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章记录未找到")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 读取成功。
		view.Render(w, view.D{
			"Articles": _article,
		}, "articles.show")
	}
}

// Edit 文章编辑页面
func (*ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 获取对应文章数据
	_article, err := article.Get(id)

	// 如果出现错误
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// w.WriteHeader(http.StatusNotFound)
			// fmt.Fprint(w, "404 文章未找到")
			view.Render(w, view.D{
				"Articles": _article,
			}, "errors.404")
		} else {
			// w.WriteHeader(http.StatusInternalServerError)
			// fmt.Fprint(w, "500 服务器内部错误")
			view.Render(w, view.D{
				"Articles": _article,
			}, "errors.50x")
		}
	} else {
		view.Render(w, view.D{
			"Articles": _article,
		}, "articles.edit", "articles._form_slide", "articles._form_field")
	}
}

// Update 文章修改
func (*ArticlesController) Update(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 获取对应文章数据
	_article, err := article.Get(id)
	// 判断是否有错误
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 未出现错误
		title := r.PostFormValue("title")
		content := r.PostFormValue("content")

		errors := validateArticleFormData(title, content)

		if len(errors) == 0 {
			_article.Title = title
			_article.Content = content

			rowsAffected, err := _article.Update()

			if err != nil {
				// 数据库错误
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
				return
			}

			// 表示更新成功
			if rowsAffected > 0 {
				showURL := route.Name2URL("articles.show", "id", id)
				http.Redirect(w, r, showURL, http.StatusFound)
			}
		} else {
			// 表示表单验证不通过 ，显示错误信息
			updateURL := route.Name2URL("articles.update", "id", id)
			data := ArticlesFormData{
				Title:   _article.Title,
				Content: _article.Content,
				URL:     updateURL,
				Errors:  errors,
			}
			view.Render(w, view.D{
				"Articles": data,
			}, "articles.edit")
		}
	}
}

// Delete 文章删除
func (*ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 读取对应的文章数据
	_article, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 未出现错误，进行删除操作
		rowsAffected, err := _article.Delete()
		// 是否发生错误
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		} else {
			// 未发生错误
			if rowsAffected > 0 {
				// 删除成功，跳转
				indexURL := route.Name2URL("articles.index")
				http.Redirect(w, r, indexURL, http.StatusFound)
			} else {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "404 文章未找到")
			}
		}
	}
}

// 验证文章表单字段
func validateArticleFormData(title string, content string) map[string]string {
	errors := make(map[string]string)
	// 验证标题
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(content) > 50 {
		errors["title"] = "标题长度应该介于 3-50 个字"
	}

	if content == "" {
		errors["content"] = "内容不能为空"
	} else if utf8.RuneCountInString(content) < 10 {
		errors["content"] = "内容长度应大于 10 个字"
	}

	return errors
}
