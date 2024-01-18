package main
import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func sayHello(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello world!",
    })
}


func main() {
	r := gin.Default()
    r.GET("/hello", sayHello)
    r.Run(":9090")
}