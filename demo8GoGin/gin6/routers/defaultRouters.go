package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "default/index.html", gin.H{
				"msg": "I am a message",
			})
		})

		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(http.StatusOK, "新闻")
		})
	}
}
