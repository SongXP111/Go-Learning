package main

import (
	"fmt"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 中间件
func initMiddlewareOne(c *gin.Context) {
	fmt.Println("11111")
	// 调用该请求的剩余处理程序
	c.Next()

	// 中止该请求的胜率处理程序
	// c.Abort()
	fmt.Println("22222")
}

// 中间件
func initMiddlewareTwo(c *gin.Context) {
	fmt.Println("33333")
	// 调用该请求的剩余处理程序
	c.Next()

	// 中止该请求的胜率处理程序
	// c.Abort()
	fmt.Println("44444")
}

func main() {
	// 创建一个默认的路由引擎，自带logger和Recovery，如果不想创建默认的可以使用New
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

	// 全局中间件
	r.Use(initMiddlewareOne, initMiddlewareTwo)

	r.GET("/", func(c *gin.Context) {
		fmt.Println("这是一个首页")
		// time.Sleep(time.Second)
		c.String(200, "gin首页")
	})

	// r.GET("/", initMiddlewareOne, initMiddlewareTwo, func(c *gin.Context) {
	// 	fmt.Println("这是一个首页")
	// 	// time.Sleep(time.Second)
	// 	c.String(200, "gin首页")
	// })

	r.GET("/news", func(c *gin.Context) {
		c.String(200, "gin新闻")
	})

	r.Run()
}
