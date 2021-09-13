package main

import (
	"io"
	"net/http"
)

type a struct{
	//id int
}

func (*a) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String() //获得访问的路径
	io.WriteString(w, path)//http://localhost:8005/...
}

func main() {
	http.ListenAndServe(":8005", &a{})//第2个参数需要实现Hander接口的struct，a满足
}