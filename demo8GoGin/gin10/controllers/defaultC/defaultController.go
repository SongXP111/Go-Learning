package defaultc

import (
	"fmt"
	"gin10/modals"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	fmt.Println(modals.UnixToTime(1629788564))
	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg": "I am a message",
		"t":   1629788418,
	})
}

func (con DefaultController) News(c *gin.Context) {
	c.String(http.StatusOK, "News")
}
