package main

import (
	"log"
	"net/http"
	"text/template"

	restful "github.com/emicklei/go-restful/v3"
)

// This example shows how to serve a HTML page using the standard Go template engine.
//
// GET http://localhost:8081/

func main() {
	ws := new(restful.WebService)//新建服务
	ws.Route(ws.GET("/").To(home))//路由路径
	restful.Add(ws)//增加服务到默认容器
	print("open browser on http://localhost:8081/\n")
	log.Fatal(http.ListenAndServe(":8081", nil))//监听
}
//文字内容

type Message struct {
	Text string
}

func home(req *restful.Request, resp *restful.Response) {
	p := &Message{"restful-html-template demo"}//信息内容
	// you might want to cache compiled templates
	t, err := template.ParseFiles("Home.html")
	if err != nil {
		log.Fatalf("Template gave: %s", err)
	}
	t.Execute(resp.ResponseWriter, p)//响应
}
