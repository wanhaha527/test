package main

import (
	
	restful "github.com/emicklei/go-restful/v3"
	"io"
	"log"
	"net/http"
)

func main() {
	ws:=new(restful.WebService)
	ws.Route(ws.GET("/test").Filter(basicAuthenticate).To(mytest))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8012",nil))
}
func basicAuthenticate(req *restful.Request,resp *restful.Response,chain *restful.FilterChain)  {
	u, p, ok :=req.Request.BasicAuth()
	if u!="admin"||p!="admin"||!ok {
		resp.AddHeader("www-Authenticate","Basic realm=my")//
		resp.WriteErrorString(401,"401:Not Authenticated")
		return
	}
	chain.ProcessFilter(req,resp)
}
func mytest (req *restful.Request, resp *restful.Response)  {
	io.WriteString(resp,"登录成功")
}