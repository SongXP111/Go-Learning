package admin

import (
	"fmt"
	"gin12/modals"
	"net/http"
	"os"
	"path"
	"strconv"

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
	username := c.PostForm("username")
	// 1. 获取上传的文件
	file, err := c.FormFile("face")
	if err == nil {
		// 2. 获取后缀名，判断类型是否正确（图片后缀）
		extName := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}
		if _, ok := allowExtMap[extName]; !ok { // 后缀名不正确
			c.String(http.StatusOK, "上传类型不合法")
			return
		}

		// 3. 创建图片保存目录
		day := modals.GetDay()
		dir := "./static/upload/" + day

		err := os.MkdirAll(dir, 0666)
		if err != nil {
			c.String(http.StatusOK, "MkdirAll失败")
			return
		}

		// 4. 生成文件名称和保存目录
		fileName := strconv.FormatInt(modals.GetUnix(), 10) + extName
		fmt.Println(fileName)

		// 5. 执行上传
		dst := path.Join(dir, fileName)
		fmt.Println(dst)
		saved := c.SaveUploadedFile(file, dst) // 保存文件（因权限无法保存）
		fmt.Println(saved)
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
