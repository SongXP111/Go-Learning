package main

import (
	"gin10/modals"
	"gin10/routers"
	"html/template"

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

	// 抽离路由配置
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)

	r.Run(":8080")
}
