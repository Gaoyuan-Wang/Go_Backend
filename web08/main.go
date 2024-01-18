package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		// 方式一：自己拼接 JSON
		data := gin.H{
			"name":    "小王子",
			"message": "hello world",
			"age":     18,
		}

		c.JSON(http.StatusOK, data)
	})
	// 方式二：使用结构体, 灵活使用tag来对结构体字段做定制化操作
	type msg struct {
		Name    string `json:"name"`
		Age     int
		Message string
	}

	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			"小王子",
			18,
			"hello golang",
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":9090")
}
