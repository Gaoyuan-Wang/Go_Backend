package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		//获取浏览器那边发请求携带的query string参数
		//name := c.DefaultQuery("name", "somebody")
		name := c.Query("name")
		age := c.Query("age")
		//name, ok := c.GetQuery("query")
		//if !ok {
		//	name = "somebody"
		//}

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":9090")
}
