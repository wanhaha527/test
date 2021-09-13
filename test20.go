package main

import (
	"github.com/emicklei/go-restful/v3"
	"io"
	"log"
	"net/http"
)

type person struct {
	id int
	name string
	age int
}


func (p person) getall (req *restful.Request,resp *restful.Response){

	p.name="www"

	io.WriteString(resp,p.name)
}
func (p person) registerto() {
	ws:=new(restful.WebService)
	ws.Path("/user")
	ws.Route(ws.GET("/{ID}}").To(p.getall))
	restful.Add(ws)

}
func main() {
person{}.registerto()
log.Fatal(http.ListenAndServe(":8020",nil))
}
