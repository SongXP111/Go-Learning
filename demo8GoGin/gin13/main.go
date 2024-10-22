package main

import (
	"gin13/modals"
	"gin13/routers"
	"html/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 自定义模板函数
	// 必须放在加载模版之前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": modals.UnixToTime,
	})
	// 加载模版
	r.LoadHTMLGlob("templates/**/*")

	// 配置静态web目录 静态服务
	r.Static("/static", "./static")

	// 配置session中间件
	// 创建基于cookie的存储引擎
	store := cookie.NewStore([]byte("secret"))
	// 配置sesion的中间件
	r.Use(sessions.Sessions("mysession", store))

	// 抽离路由配置
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)

	r.Run(":8080")
}
