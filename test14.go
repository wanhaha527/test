package main

import (
	restful "github.com/emicklei/go-restful/v3"
	"io"
	"log"
	"net/http"
)

type UserResource struct {}

func (u UserResource) RegisterTo(container *restful.Container)  {
	ws:=new(restful.WebService)
	ws.Path("/users").
		Consumes("*/*").
		Produces("*/*")
	ws.Route(ws.GET("/{user-id}}").To(u.nop))
	ws.Route(ws.POST("").To(u.nop))
	ws.Route(ws.PUT("/{user-id}").To(u.nop))
	ws.Route(ws.DELETE("/{user-id}}").To(u.nop))
	container.Add(ws)
}
func (u UserResource) nop(rep *restful.Request,resp *restful.Response)  {
	io.WriteString(resp.ResponseWriter,"This is a test")
}
func main() {

	wscontainer := restful.NewContainer()
	u:=UserResource{}
	u.RegisterTo(wscontainer)//资源注册到容器

	cors:=restful.CrossOriginResourceSharing{
		ExposeHeaders:         []string{"X-My_Header"},
		AllowedHeaders:        []string{"Content-Type","Accept"},
		AllowedMethods:        []string{"GET","POST"},
		CookiesAllowed:        false,
		Container:             wscontainer}
	wscontainer.Filter(cors.Filter)
	wscontainer.Filter(wscontainer.OPTIONSFilter)
	log.Printf("Starting listening on port:8014")
	log.Fatal(http.ListenAndServe(":8014",wscontainer))

}
