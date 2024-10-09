package main

import (
	"gin6/routers"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	r := gin.Default()
	// 自定义模板函数
	// 必须放在加载模版之前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
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
