package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/go",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.RemoteAddr, "链接成功")
			//请求方式：GET POST DELETE PUT UPDATE
			fmt.Println("method", r.Method)
			// /go
			fmt.Println("url", r.URL.Path)
			fmt.Println("header:", r.Header)
			fmt.Println("body", r.Body)
			//
			w.Write([]byte("www.baidu.com"))
		})
	    http.ListenAndServe("127.0.0.1:8080", nil)
}
