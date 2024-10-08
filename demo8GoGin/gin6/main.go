package main

import (
	"html/template"
	"net/http"
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

	// 静态服务
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "首页")
	})

	r.Run(":8080")
}
