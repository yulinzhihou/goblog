// routes
// Author : Yulinzhihou
// GitHub : https://github.com/yulinzhihou
// WebSite: yulinzhihou.com
// Date   : 2023/12/21

package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"myblog/app/http/controllers"
	"myblog/app/http/middlewares"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {
	// 静态页面
	pc := new(controllers.PagesController)
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")
	r.HandleFunc("/help", pc.Help).Methods("GET").Name("help")

	// 文章相关页面
	ac := new(controllers.ArticlesController)

	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")
	r.HandleFunc("/articles/create", middlewares.Auth(ac.Create)).Methods("GET").Name("articles.create")
	r.HandleFunc("/articles", middlewares.Auth(ac.Store)).Methods("POST").Name("articles.store")
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	r.HandleFunc("/articles/{id:[0-9]+}", middlewares.Auth(ac.Update)).Methods("POST").Name("articles.update")
	r.HandleFunc("/articles/{id:[0-9]+}/edit", middlewares.Auth(ac.Edit)).Methods("GET").Name("articles.edit")
	r.HandleFunc("/articles/{id:[0-9]+}/delete", middlewares.Auth(ac.Delete)).Methods("POST").Name("articles.delete")

	// 注册登录相送
	auc := new(controllers.AuthController)
	r.HandleFunc("/auth/register", middlewares.Guest(auc.Register)).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do_register", middlewares.Guest(auc.DoRegister)).Methods("POST").Name("auth.do_register")
	r.HandleFunc("/auth/login", middlewares.Guest(auc.Login)).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/do_login", middlewares.Guest(auc.DoLogin)).Methods("POST").Name("auth.do_login")
	r.HandleFunc("/auth/forget", middlewares.Guest(auc.Forget)).Methods("GET").Name("auth.forget")
	r.HandleFunc("/auth/reset", middlewares.Guest(auc.ResetPassword)).Methods("POST").Name("auth.reset")
	r.HandleFunc("/auth/send_email", middlewares.Guest(auc.SendEmail)).Methods("GET").Name("auth.send_email")
	r.HandleFunc("/auth/do_send_email", middlewares.Guest(auc.DoSendEmail)).Methods("POST").Name("auth.do_send_email")
	r.HandleFunc("/auth/logout", middlewares.Auth(auc.Logout)).Methods("POST").Name("auth.logout")

	// 用户相关
	uc := new(controllers.UserController)
	r.HandleFunc("/users/{id:[0-9]+}", middlewares.Auth(uc.Show)).Methods("GET").Name("users.show")
	r.HandleFunc("/users/{id:[0-9]+}", middlewares.Auth(uc.Update)).Methods("POST").Name("users.update")

	// 分类相关
	cc := new(controllers.CategoriesController)
	r.HandleFunc("/categories/create", middlewares.Auth(cc.Create)).Methods("GET").Name("categories.create")
	r.HandleFunc("/categories", middlewares.Auth(cc.Store)).Methods("POST").Name("categories.store")
	r.HandleFunc("/categories", pc.NotFound).Methods("GET")
	r.HandleFunc("/categories/{id:[0-9]+}", middlewares.Auth(cc.Show)).Methods("GET").Name("categories.show")

	// 静态资源
	r.PathPrefix("/assets/").Handler(http.FileServer(http.Dir("./public")))

	// 使用中间件
	r.Use(middlewares.StartSession)
}
