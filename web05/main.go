package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./index2.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	//渲染模板
	msg := "小王子"
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println("render template failed, err:%v", err)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./home2.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	//渲染模板
	msg := "小王子"
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println("render template failed, err:%v", err)
		return
	}
}

func index2(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./template/base.tmpl", "./template/index2.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	//渲染模板
	msg := "小王子"
	err = t.ExecuteTemplate(w, "index2.tmpl", msg)
	if err != nil {
		fmt.Println("render template failed, err:%v", err)
		return
	}
}

func home2(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./template/base.tmpl", "./template/home2.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	//渲染模板
	msg := "七米"
	err = t.ExecuteTemplate(w, "home2.tmpl", msg)
	if err != nil {
		fmt.Println("render template failed, err:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server failed, err:%v", err)
		return
	}
}
