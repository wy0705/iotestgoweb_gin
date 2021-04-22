package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func VIII8() {
	r := gin.Default()
	r.LoadHTMLGlob("/disk/go15pro/src/iotestgoweb_gin/web/viii.html")

	r.GET("/index",func(c *gin.Context){
		c.HTML(http.StatusOK, "viii.html",nil)
	})
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// r.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// 单个文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Println(file.Filename)
		dst := fmt.Sprintf("/disk/go15pro/src/iotestgoweb_gin/web/upload/%s", file.Filename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})

	r.POST("/upload2", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["f1"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("/disk/go15pro/src/iotestgoweb_gin/web/upload/%d_%s", index,file.Filename)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	r.Run("0.0.0.0:9999")
}
