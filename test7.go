package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type User struct {
	Id ,Name string
}
type UserResoure struct {
	users map[string] User
}
func (u UserResoure) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON,restful.MIME_XML)
	ws.Route(ws.POST("/{user-id}").To(u.findUser))
	//ws.Route(ws.POST("").To(u.updateUser))
	//ws.Route(ws.PUT("/{user-id}").To(u.createUser))
	//ws.Route(ws.PUT("/{user-id}").To(u.removeUser))
	container.Add(ws)
}
func (u UserResoure) findUser(request *restful.Request , response *restful.Response) {
	id := request.PathParameter("user-id")
	usr , ok := u.users[id]
	if !ok {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound , "users could not be found.")
	} else {
		response.WriteEntity(usr)
	}
}
/*
func (u *UserResource) updateUser(request *restful.Request, response *restful.Response) { }
func (u *UserResource) createUser(request *restful.Request, response *restful.Response) { }
func (u *UserResource) removeUser(request *restful.Request, response *restful.Response) { }
*/
func main() {
	wsContainer := restful.NewContainer ()
	wsContainer.Router(restful.CurlyRouter{})
	u :=UserResoure{map[string]User{}}
	u.Register(wsContainer)

	log.Printf("start listening on localhost:8007")
	server := &http.Server{Addr: ":8007", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}


