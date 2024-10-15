package routers

import (
	"gin10/controllers/admin"
	"gin10/middlewares"

	"github.com/gin-gonic/gin"
)

// 首字母大写代表公有方法
func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middlewares.InitMiddleware)
	{
		adminRouters.GET("/", admin.IndexController{}.Index)

		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.POST("/user/doUpload", admin.UserController{}.DoUpload)
		adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		adminRouters.POST("/user/doEdit", admin.UserController{}.DoEdit)

		adminRouters.GET("/article", admin.ArticleController{}.Index)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/edit", admin.ArticleController{}.Edit)
	}
}
