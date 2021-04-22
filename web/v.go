package web

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func V5() {
	r := gin.Default()

	// gin.H 是map[string]interface{}的缩写
	r.GET("/json1", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	r.GET("/json2", func(c *gin.Context) {
		// 方法二：使用结构体
		var msg struct {
			Name    string `json:"user"` //可以使用tag来修改json的key
			Sex string
			Age     int
		}
		msg.Name = "xxx"
		msg.Sex = "n"
		msg.Age = 22
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/xml1", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.XML(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	r.GET("/xml2", func(c *gin.Context) {
		// 方法二：使用结构体
		type MessageRecord struct {
			Name    string
			Sex string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "xxx"
		msg.Sex = "n"
		msg.Age = 22
		c.XML(http.StatusOK, msg)
	})

	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	})

	r.GET("/protobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})
	r.Run("0.0.0.0:9999")
}
