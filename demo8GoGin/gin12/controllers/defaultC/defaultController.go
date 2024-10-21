package defaultc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	// 设置cookie
	c.SetCookie("username", "zhangsan", 3600, "/", "localhost", false, false)
	// 过期时间演示
	c.SetCookie("hobby", "吃饭", 5, "/", "localhost", false, false)
	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg": "I am a message",
		"t":   1629788418,
	})
}

func (con DefaultController) News(c *gin.Context) {
	// 获取cookie
	username, _ := c.Cookie("username")
	hobby, _ := c.Cookie("hobby")
	c.String(http.StatusOK, "Cookie: username: "+username+", hobby: "+hobby)
}

func (con DefaultController) Shop(c *gin.Context) {
	// 获取cookie
	username, _ := c.Cookie("username")
	hobby, _ := c.Cookie("hobby")
	c.String(http.StatusOK, "Cookie: username: "+username+", hobby: "+hobby)
}

func (con DefaultController) DeleteCookie(c *gin.Context) {
	c.SetCookie("username", "zhangsan", -1, "/", "localhost", false, false)
	c.String(http.StatusOK, "删除cookie成功")
}
