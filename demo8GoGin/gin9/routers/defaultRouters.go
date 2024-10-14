package routers

import (
	defaultc "gin9/controllers/defaultC"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", defaultc.DefaultController{}.Index)

		defaultRouters.GET("/news", defaultc.DefaultController{}.News)
	}
}
