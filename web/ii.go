package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func II2() {
	r := gin.Default()
	r.LoadHTMLGlob("/disk/go15pro/src/iotestgoweb_gin/web/ii.tmpl")
	//r.LoadHTMLFiles("/disk/go15pro/src/iotestgoweb_gin/web/ii.tmpl", "/disk/go15pro/src/iotestgoweb_gin/web/ii2.tmpl")
	r.GET("/xxx/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ii/index.html", gin.H{
			"name": "aaabbb",
		})
	})

	r.Run("0.0.0.0:9999")
}
