package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func VI6() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.LoadHTMLGlob("/disk/go15pro/src/iotestgoweb_gin/web/vi.html")

	r.GET("/index",func(c *gin.Context){
		c.HTML(http.StatusOK, "vi.html",nil)
	})

	r.GET("/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "xxx")
		//username := c.Query("username")
		password := c.Query("password")
		xxx:=c.QueryArray("xxx")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"password":  password,
			"xxx":xxx,
		})
	})

	r.GET("/search/:username/:password", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"password":  password,
		})
	})

	r.POST("/search", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		//username := c.DefaultPostForm("username", "小王子")
		username := c.PostForm("username")
		password := c.PostForm("password")
		xxx:=c.PostFormArray("xxx")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"password":  password,
			"xxx": xxx,
		})
	})

	r.Run("0.0.0.0:9999")
}
