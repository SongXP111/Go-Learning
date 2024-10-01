package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*") // 配置模版的文件
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Value: %v", "HomePage")
	})
	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "Hello Gin1",
		})
	})
	r.GET("/json2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "Hello Gin2",
		})
	})
	r.GET("/json3", func(c *gin.Context) {
		a := &Article{
			Title:   "Title",
			Desc:    "Desc",
			Content: "Content",
		}
		c.JSON(200, a)
	})

	// 响应JSONP请求
	r.GET("/jsonp", func(c *gin.Context) {
		a := &Article{
			Title:   "Title",
			Desc:    "Desc",
			Content: "Content",
		}
		c.JSONP(200, a)
	})

	// XML
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "Hello GinXML",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "Title",
		})
	})

	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "Goods Page",
			"price": 20,
		})
	})

	r.Run(":8080")
}
