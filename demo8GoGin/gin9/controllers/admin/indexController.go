package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (con IndexController) Index(c *gin.Context) {
	username, _ := c.Get("username") // 通过中间件传输数据
	fmt.Println(username)

	v, ok := username.(string)
	if ok {
		c.String(http.StatusOK, "用户列表 --- "+v)
	} else {
		c.String(http.StatusOK, "用户列表 --- "+"获取用户失败")
	}
}
