package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模版
	t, err := template.ParseFiles("./go.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	// 渲染模版
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("Render template failed, err:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err:%v", err)
		return
	}
}
