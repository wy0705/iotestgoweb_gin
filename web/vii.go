package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
type Login struct {
	Username     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
func VII7() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.LoadHTMLGlob("/disk/go15pro/src/iotestgoweb_gin/web/vii.html")

	r.GET("/index",func(c *gin.Context){
		c.HTML(http.StatusOK, "vii.html",nil)
	})

	// 绑定JSON的示例 ({"username": "aaa", "password": "bbb"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var login Login

		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"username":     login.Username,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定form表单示例 (user=aaa&password=bbb)
	r.POST("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.Username,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定QueryString示例 (/loginQuery?user=aaa&password=bbb)
	r.GET("/loginGet", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.Username,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run("0.0.0.0:9999")

}
