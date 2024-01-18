package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.tmpl").Delims("{[", "]}").ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	msg := "Hello World"
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println("Execute template failed, err:%v", err)
		return
	}
}

func xss(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	str1 := "<script>alert(123)</script>"
	str2 := "<a href='https://www.baidu.com'>百度</a>"
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server failed, err:%v", err)
		return
	}
}
