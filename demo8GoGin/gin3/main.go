package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string
	Content string
}

func main() {
	r := gin.Default()
	// 加载模版
	r.LoadHTMLGlob("templates/**/*")

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
