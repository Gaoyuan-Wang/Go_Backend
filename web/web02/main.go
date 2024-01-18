package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	// 渲染模版
	u1 := User{ // u1.Name
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	m1 := map[string]interface{}{ // m1.Name
		"Name":   "小王子",
		"Gender": "男",
		"Age":    18,
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}
	t.Execute(w, map[string]interface{}{
		"u1":        u1,
		"m1":        m1,
		"hobbyList": hobbyList,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err:%v", err)
		return
	}
}
