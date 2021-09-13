package main

import (

	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type User struct {
	Id   string
	Name string
}
type UserResource struct {
	// normally one would use DAO (data access object)
	users map[string]User
}

func (u UserResource) Register(container *restful.Container){
	ws:=new(restful.WebService)
	ws.Path("/users").Consumes(restful.MIME_XML,restful.MIME_JSON).Produces(
		restful.MIME_JSON,restful.MIME_XML)//Consumes指定处理请求的提交内容类型,Produces指定返回的内容类型
	ws.Route(ws.GET("/{user-id}").To(u.findUser))
	ws.Route(ws.POST("").To(u.updateUser))
	ws.Route(ws.PUT("/{user-id}").To(u.creatUser))
	ws.Route(ws.DELETE("/{user-id}").To(u.removeUser))
	container.Add(ws)

}
func (u UserResource) findUser(req *restful.Request,resp *restful.Response){
	id:=req.PathParameter("user-id")
	usr,ok:=u.users[id]
	if !ok {
		resp.AddHeader("Content-Type","text/plain")
		resp.WriteErrorString(http.StatusNotFound,"User could not be found")
	}else{
		resp.WriteEntity(usr)
	}
}

func (u UserResource) updateUser(req *restful.Request,resp *restful.Response){
	usr:=new(User)
	err:=req.ReadEntity(&usr)
	if err==nil {
		u.users[usr.Id]=*usr
		resp.WriteEntity(usr)
	}else{
		resp.AddHeader("Content-Type","text/plain")
		resp.WriteErrorString(http.StatusInternalServerError,err.Error())
	}
}

func (u UserResource) creatUser(req *restful.Request,resp *restful.Response)  {
	usr:=User{Id:req.PathParameter("user-id")}
	err:=req.ReadEntity(&usr)
	if err==nil {
		u.users[usr.Id]=usr
		resp.WriteHeaderAndEntity(http.StatusCreated,usr)
	}else {
		resp.AddHeader("Content-Type","text/plain")
		resp.WriteErrorString(http.StatusInternalServerError,err.Error())
	}
}

func (u UserResource) removeUser(req *restful.Request,resp *restful.Response)  {
	id:=req.PathParameter("user-id")
	delete(u.users,id)
}

func main() {
	wsc:=restful.NewContainer()//新建服务器
	wsc.Router(restful.CurlyRouter{})//路由分发器
	u:=UserResource{map[string]User{}}
	u.Register(wsc)//将 WebServices 关联到 Container

	log.Printf("Start listening on port 8010")
	//server:=&http.Server{Addr: ":8010", Handler: wsc}
	log.Fatal(http.ListenAndServe(":8010",wsc))

}
