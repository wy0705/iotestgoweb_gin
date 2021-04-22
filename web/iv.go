package web

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func IV4() {
	r := gin.Default()
	r.HTMLRender = xxx("/disk/go15pro/src/iotestgoweb_gin/web/iv_tmpl")
	r.GET("/index", func (c *gin.Context){
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	r.GET("/home", func (c *gin.Context){
		c.HTML(http.StatusOK, "home.tmpl", nil)
	})

	r.Run("0.0.0.0:9999")
}
func xxx(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	// 为layouts/和includes/目录生成 templates map
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
