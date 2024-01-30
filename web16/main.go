package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Println("...index")
	name, ok := c.Get("name")
	if !ok {
		name = "default name"
	}
	c.JSON(http.StatusOK, gin.H{
		"message": name,
	})
}

func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	start := time.Now()
	c.Next() //调用后续的处理函数
	//c.Abort()//阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out ...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Set("name", "gaoyuan")
	c.Next() //调用后续的处理函数
	//c.Abort()//阻止调用后续的处理函数
	fmt.Println("m2 out ...")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if doCheck {
			//c.Next()
			//c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(m1, m2, authMiddleware(true))
	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "shop",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "user",
		})
	})
	r.Run(":8080")
}
