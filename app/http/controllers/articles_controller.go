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

	"gorm.io/gorm"
	"myblog/app/models/article"
	"myblog/pkg/logger"
	"myblog/pkg/route"
)

// ArticlesController 文章相关页面
type ArticlesController struct {
}

type ArticlesFormData struct {
	ID      uint64
	Image   string
	Brief   string
	Content string
	Title   string
	Status  int8
	UserID  uint64
	URL     string
	Errors  map[string]string
}

// Index 文章列表页
func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	// 获取结果集
	articles, err := article.GetAll()

	if err != nil {
		// 数据库错误
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	} else {
		// 加载模板
		tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
		logger.LogError(err)

		// 渲染模板
		err = tmpl.Execute(w, articles)
		logger.LogError(err)
	}
}

// Create 文章创建表单页面
func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	storeURL := route.Name2URL("articles.store")

	data := ArticlesFormData{
		Image:  "",
		URL:    storeURL,
		Errors: nil,
	}
	tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, data)
	logger.LogError(err)
}

// Store 文章新增
func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	// title := r.PostFormValue("title")
	// content := r.PostFormValue("content")
	//
	// errors := validateArticleFormData(title, content)
	//
	// // 检查是否有错误
	// if len(errors) == 0 {
	// 	lastInsertID, err := saveArticleToDB(title, content)
	// 	if lastInsertID > 0 {
	// 		fmt.Fprint(w, "插入成功，ID 为"+strconv.FormatUint(lastInsertID, 10))
	// 	} else {
	// 		logger.LogError(err)
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		fmt.Fprint(w, "500 服务器内部错误")
	// 	}
	// } else {
	// 	storeURL := route.Name2URL("articles.store")
	//
	// 	data := ArticlesFormData{
	// 		Title:   title,
	// 		Content: content,
	// 		URL:     storeURL,
	// 		Errors:  errors,
	// 	}
	//
	// 	tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
	// 	logger.LogError(err)
	//
	// 	err = tmpl.Execute(w, data)
	// 	logger.LogError(err)
	// }
}

// Show 文章详情页面
func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {

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
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 读取成功
		updateURL := route.Name2URL("articles.update", "id", id)

		data := ArticlesFormData{
			Title:   _article.Title,
			Content: _article.Content,
			URL:     updateURL,
			Errors:  nil,
		}

		tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
		logger.LogError(err)

		err = tmpl.Execute(w, data)
		logger.LogError(err)
	}
}

// Update 文章修改
func (*ArticlesController) Update(w http.ResponseWriter, r *http.Request) {
	// // 获取 URL 参数
	// id := route.GetRouteVariable("id", r)
	// // 获取对应文章数据
	// _article, err := article.Get(id)
	// // 判断是否有错误
	// if err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		w.WriteHeader(http.StatusNotFound)
	// 		fmt.Fprint(w, "404 文章未找到")
	// 	} else {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		fmt.Fprint(w, "500 服务器内部错误")
	// 	}
	// } else {
	// 	// 未出现错误
	// 	title := r.PostFormValue("title")
	// 	content := r.PostFormValue("content")
	//
	// 	errors := validateArticleFormData(title, content)
	//
	// 	if len(errors) == 0 {
	// 		_article.Title = title
	// 		_article.Content = content
	//
	// 		rowsAffected, err := _article.Update()
	//
	// 		if err != nil {
	// 			// 数据库错误
	// 			w.WriteHeader(http.StatusInternalServerError)
	// 			fmt.Fprint(w, "500 服务器内部错误")
	// 			return
	// 		}
	//
	// 		// 表示更新成功
	// 		if rowsAffected > 0 {
	// 			showURL := route.Name2URL("articles.show", "id", id)
	// 			http.Redirect(w, r, showURL, http.StatusFound)
	// 		}
	// 	} else {
	// 		// 表示表单验证不通过 ，显示错误信息
	// 		updateURL := route.Name2URL("articles.update", "id", id)
	// 		data := ArticlesFormData{
	// 			Title:   _article.Title,
	// 			Content: _article.Content,
	// 			URL:     updateURL,
	// 			Errors:  errors,
	// 		}
	// 		tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
	// 		logger.LogError(err)
	//
	// 		err = tmpl.Execute(w, data)
	// 		logger.LogError(err)
	// 	}
	// }
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
