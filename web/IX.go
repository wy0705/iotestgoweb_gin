package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IX9() {
	r := gin.Default()
	r.LoadHTMLGlob("/disk/go15pro/src/iotestgoweb_gin/web/404.html")
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	r.GET("/test1", func(c *gin.Context) {
		// 指定重定向的URL
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})

	//get post delete put等均可
	r.Any("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	//组路由
	test3Group := r.Group("/test3")
	{
		test3Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"hello": "test3index"})
		})
		test3Group.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"hello": "test3home"})
		})

	}
	test4Group := r.Group("/test4")
	{
		test4Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"hello": "test4index"})
		})
		test4Group.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"hello": "test4home"})
		})
	}

	//组路由嵌套
	test5Group := r.Group("/test5")
	{
		// 嵌套路由组
		xx := test5Group.Group("aaa")
		xx.GET("/bbb", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"hello": "test5"})
		})
	}

	r.Run("0.0.0.0:9999")
}
