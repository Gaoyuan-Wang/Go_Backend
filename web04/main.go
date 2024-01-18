package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数 kua
	k := func(name string) (string, error) {
		return name + "年轻又帅气", nil
	}
	// 定义模板
	t := template.New("f.tmpl")
	// 告诉模版对象，我现在多了一个自定义的函数 kua
	t.Funcs(template.FuncMap{
		"kua": k,
	})
	// 解析模板
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:", err)
		return
	}
	name := "小王子"
	// 渲染模板
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("Execute template failed, err:%v\n", err)
		return
	}
}

func demo1(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v\n", err)
		return
	}
	name := "小王子"
	// 渲染模板
	err = t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server failed, err:", err)
		return
	}
}
