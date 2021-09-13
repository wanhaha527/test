package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"path"
)
var rootdir=""
func main() {
	restful.DefaultContainer.Router(restful.CurlyRouter{})
	ws:=new(restful.WebService)
	ws.Route(ws.GET("/static/{subpath:*}").To(StaticFromPathParam))
	ws.Route(ws.GET("/static").To(StaticFromQueryPrarm))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8017",nil))
}
func StaticFromPathParam(req *restful.Request,resp *restful.Response)  {
	actul:=path.Join(rootdir,req.PathParameter("subpath"))
	fmt.Printf("serving %s from %s \n",actul,req.PathParameter("subpath"))
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		actul)

}
func StaticFromQueryPrarm(req *restful.Request,resp *restful.Response)  {
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		path.Join(rootdir,req.QueryParameter("resource")))
}