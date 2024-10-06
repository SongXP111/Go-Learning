package main

import (
	"fmt"
	"html/template"
	"net/http"
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

func Println(str1, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + str2
}

func main() {
	r := gin.Default()
	// 自定义模板函数
	// 必须放在加载模版之前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})
	// 加载模版
	r.LoadHTMLGlob("templates/**/*")

	// 静态服务
	r.Static("/static", "./static")

	// 前台
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
			"msg":   "我是msg",
			"score": 91,
			"hobby": []string{"吃饭", "睡觉", "玩游戏", "写代码"},
			"newsList": []interface{}{
				&Article{
					Title:   "新闻标题111",
					Content: "新闻内容111",
				},
				&Article{
					Title:   "新闻标题222",
					Content: "新闻内容222",
				},
			},
			"testSlice": []string{},
			"news": &Article{
				Title:   "新闻标题",
				Content: "新闻内容",
			},
			"date": 1629423555,
		})
	})
	r.GET("/news", func(c *gin.Context) {
		a := &Article{
			Title:   "新闻标题",
			Content: "新闻内容",
		}
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻页面",
			"news":  a,
		})
	})
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})
	r.Run(":8080")
}
