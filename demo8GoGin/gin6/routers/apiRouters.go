package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "我是一个API接口")
		})

		apiRouters.GET("/userlist", func(c *gin.Context) {
			c.String(http.StatusOK, "我是一个API接口-userlist")
		})

		apiRouters.GET("/plist", func(c *gin.Context) {
			c.String(http.StatusOK, "我是一个API接口-plist")
		})

		apiRouters.GET("/cart", func(c *gin.Context) {
			c.String(http.StatusOK, "我是一个API接口")
		})
	}
}
