package main

import (
	"fmt"
	"net/http"
)

// http server

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World！")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8070", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
