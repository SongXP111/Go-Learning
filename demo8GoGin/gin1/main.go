package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Value：%v", "Hello Gin")
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(http.StatusOK, "I am news 123")
	})
	r.POST("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "I am a post request")
	})
	r.PUT("/edit", func(c *gin.Context) {
		c.String(http.StatusOK, "I am a put request")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(http.StatusOK, "I am a delete request")
	})
	r.Run(":8080") // 在8080端口启动一个web服务
}
