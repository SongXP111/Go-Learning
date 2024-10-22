package routers

import (
	defaultc "gin13/controllers/defaultC"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", defaultc.DefaultController{}.Index)
		defaultRouters.GET("/news", defaultc.DefaultController{}.News)
		defaultRouters.GET("/shop", defaultc.DefaultController{}.Shop)
		defaultRouters.GET("/deleteCookie", defaultc.DefaultController{}.DeleteCookie)
	}
}
