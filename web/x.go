package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func X10() {
	/**
	  所有的接口都要由路由来进行管理。
	      Gin的路由支持GET,POST,PUT,DELETE,PATCH,HEAD,OPTIONS等请求
	      同时还有一个Any函数，可以同时支持以上的所有请求。

	  创建路由(router)并引入默认中间件
	      router := gin.Default()
	      在源码中,首先是New一个engine,紧接着通过Use方法传入了Logger()和Recovery()这两个中间件。
	      其中 Logger 是对日志进行记录，而 Recovery 是对有 painc时, 进行500的错误处理。
	  创建路由(router)无中间件
	      router := gin.New()
	*/
	r := gin.Default()

	r.GET("/cookie", func(c *gin.Context) {
		//获取Cookie
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			//设置cookie
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": cookie,
		})
	})

	//定义加密
	store := cookie.NewStore([]byte("secret"))
	//定义加密（将session信息存储在redis服务器）
	//store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))

	//绑定session中间件
	r.Use(sessions.Sessions("mysession", store))


	//定义GET方法
	r.GET("/session", func(c *gin.Context) {
		//初始化session对象
		session := sessions.Default(c)

		//如果浏览器第一次访问返回状态码401，第二次访问则返回状态码200
		if session.Get("username") != "xxx" {
			session.Set("username", "xxx")
			session.Save()
			c.JSON(http.StatusUnauthorized, gin.H{"username": session.Get("username")})
		} else {
			c.String(http.StatusOK, "Successful second visit")
		}

	})

	r.Run("0.0.0.0:9999")

}
