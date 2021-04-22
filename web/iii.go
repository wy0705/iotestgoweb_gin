package web

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func III3() {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})

	//两个*表示目录，一个*表示模板文件
	r.LoadHTMLGlob("/disk/go15pro/src/iotestgoweb_gin/web/iii_tmpl/**/*")

	r.Static("/static", "/disk/go15pro/src/iotestgoweb_gin/web/static")

	r.GET("/iii1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iii1/index.html", gin.H{
			"name": "aaabbb",
		})
	})

	r.GET("/iii2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "iii2/index.html", gin.H{
			"name": "<script>alert(\"aaa\");</script>",
		})
	})

	r.Run("0.0.0.0:9999")
}
