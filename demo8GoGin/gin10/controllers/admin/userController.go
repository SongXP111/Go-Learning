package admin

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController // 继承
}

func (con UserController) Index(c *gin.Context) {
	con.success(c)
}

func (con UserController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
}

func (con UserController) DoUpload(c *gin.Context) {
	// 单文件
	// username := c.PostForm("username")
	// file, err := c.FormFile("face")
	// // file.Filename 获取文件名称
	// dst := path.Join("./static/upload", file.Filename)
	// if err == nil {
	// 	c.SaveUploadedFile(file, dst) // 保存文件
	// }
	// // c.String(200, "上传文件")
	// c.JSON(http.StatusOK, gin.H{
	// 	"success":  true,
	// 	"username": username,
	// 	"dst":      dst,
	// })

	// 重名文件
	username := c.PostForm("username")
	form, _ := c.MultipartForm()
	files := form.File["face[]"]

	for _, file := range files {
		dst := path.Join("./static/upload", file.Filename)
		c.SaveUploadedFile(file, dst)
	}
	// c.String(200, "上传文件")
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
	})
}

func (con UserController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useredit.html", gin.H{})
}

func (con UserController) DoEdit(c *gin.Context) {
	username := c.PostForm("username")
	file1, err1 := c.FormFile("face1")
	// file.Filename 获取文件名称
	dst1 := path.Join("./static/upload", file1.Filename)
	if err1 == nil {
		c.SaveUploadedFile(file1, dst1) // 保存文件
	}

	file2, err2 := c.FormFile("face2")
	// file.Filename 获取文件名称
	dst2 := path.Join("./static/upload", file2.Filename)
	if err2 == nil {
		c.SaveUploadedFile(file2, dst2) // 保存文件
	}
	// c.String(200, "上传文件")
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
		"dst1":     dst1,
		"dst2":     dst2,
	})
}
