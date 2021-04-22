package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 过滤器
func Myfilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "xxx") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		log.Println("in...")
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		log.Println("out...")
		// 计算耗时
		end := time.Since(start)
		log.Println(end)
	}
}
func XI11() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()
	// 注册一个全局中间件
	r.Use(Myfilter())

	r.GET("/xxx", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	// 给/test2路由单独注册中间件（可注册多个）
	r.GET("/xxx2", Myfilter(), func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	//为路由组注册中间件
	testGroup := r.Group("/xxx3", Myfilter())
	{
		testGroup.GET("/index", func(c *gin.Context) {
			name := c.MustGet("name").(string) // 从上下文取值
			log.Println(name)
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello world!",
			})
		})
	}
	//为路由组注册中间件第二种写法
	testGroup2 := r.Group("/xxx4")
	testGroup2.Use(Myfilter())
	{
		testGroup2.GET("/index", func(c *gin.Context) {
			name := c.MustGet("name").(string) // 从上下文取值
			log.Println(name)
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello world!",
			})
		})
	}
	r.Run("0.0.0.0:9999")
}
