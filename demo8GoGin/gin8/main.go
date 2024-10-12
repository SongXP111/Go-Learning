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
func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("11111")
	// 调用该请求的剩余处理程序
	c.Next()
	fmt.Println("22222")
	end := time.Now().UnixNano()
	fmt.Println(end - start)
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

	r.GET("/", initMiddleware, func(c *gin.Context) {
		fmt.Println("这是一个首页")
		time.Sleep(time.Second)
		c.String(200, "gin首页")
	})

	r.GET("/news", initMiddleware, func(c *gin.Context) {
		c.String(200, "gin新闻")
	})

	r.Run()
}
