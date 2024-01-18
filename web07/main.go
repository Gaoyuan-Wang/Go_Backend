package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	//加载静态文件
	r.Static("/xxx", "./static")
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//r.LoadHTMLFiles("template/index.tmpl") // 模板文件解析
	r.LoadHTMLGlob("template/**/*") // 模板文件解析
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ // 模板文件渲染
			"title": "liwenzhou.com",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{ // 模板文件渲染
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})
	r.Run(":9090") //启动 server
}
