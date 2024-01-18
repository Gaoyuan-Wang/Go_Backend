package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "通过自己创建的处理器处理请求")
	fmt.Fprintf(w, "通过详细配置服务器的信息来处理请求")
}

func main() {

	myHander := MyHandler{}
	//http.Handle("/myHandler", &myHander)

	server := http.Server{
		Addr: ":8080",
		Handler: &myHander,
		ReadTimeout: 2 * time.Second,
	}

	//http.ListenAndServe(":8080", nil)
	server.ListenAndServe()
}