package main

import (

	"github.com/emicklei/go-restful/v3"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)
var logger *log.Logger=log.New(os.Stdout,"",0)

 //通用日志格式化记录器
func NCSACommonLogFormatLogger() restful.FilterFunction {
	return func(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
		var username = ""
		if  req.Request.URL.User!=nil{
			if name:=req.Request.URL.User.Username(); name!="" {
				username=name
			}

		}
		chain.ProcessFilter(req,resp)
		logger.Printf("%s - %s [%s] \"%s %s %s\" %d %d",
			strings.Split(req.Request.RemoteAddr,":")[0],
			username,
			time.Now().Format("01/jan/2020:12:22:02 -0600"),
			req.Request.Method,
			req.Request.URL.RequestURI(),
			req.Request.Proto,
			resp.StatusCode(),
			resp.ContentLength(),
			)

	}


}
func main() {
	ws:=new(restful.WebService)//新建一个Webservice
	ws.Filter(NCSACommonLogFormatLogger())//过滤器，过滤器将过滤器功能添加到适用于其所有路线的过滤器链中
	ws.Route(ws.GET("/ping").To(hello))//路由路径，Route使用RouteBuilder创建新路由，并添加到路由的有序列表中。
	restful.Add(ws)//加入Webservice到默认容器
	log.Fatal(http.ListenAndServe(":8015",nil))//监听端口
}
func hello(req *restful.Request,resp *restful.Response)  {
	io.WriteString(resp,"test15:pong")
}
