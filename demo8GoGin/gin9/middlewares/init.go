package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	// 判断用户是否登陆
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)

	// 通过中间件传输数据
	c.Set("username", "zhangsan")

	// 定义一个goroutine统计日志
	cCopy := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done in path " + cCopy.Request.URL.Path)
	}()
}
