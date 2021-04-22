package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func I1() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	//restfull
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{//http包封装了状态码常量
			"message": "GET",
		})
	})

	r.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
	// 启动HTTP服务，如果不写参数,默认0.0.0.0:8080
	r.Run("0.0.0.0:9999")
}
