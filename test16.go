package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

/*此示例演示如何将方法用作Web服务的RouteFunction。
ProductResource有一个Register（）方法，用于创建和初始化
将其方法公开为REST操作的Web服务。
WebService被添加到restful.DefaultContainer中。
ProductResource通常是使用某些数据访问对象创建的。*/

type Product struct {
	Id string
	Title string
}
type ProductResource struct {

}

func (p ProductResource) getOne(req *restful.Request,resp *restful.Response)  {
	id:=req.PathParameter("ID")//路径参数名ID必须与ws.GET("/{ID}")中一致
	log.Printf("Getting product id "+id)
	//io.WriteString(resp.ResponseWriter,"ss")
	resp.WriteEntity(Product{Id:id,Title: "test"})//写入实体Product{}


}
func (p ProductResource) postOne(req *restful.Request,resp *restful.Response)  {
	updatedProduct:=new(Product)
	err:=req.ReadEntity(updatedProduct)
	if err != nil {
		resp.WriteErrorString(http.StatusBadRequest,err.Error())
		return
	}
	log.Printf("Updating product id"+updatedProduct.Id)
}
func (p ProductResource) Register()  {
	ws:=new(restful.WebService)//新建Webservice
	ws.Path("/products").Consumes(restful.MIME_XML).Produces(restful.MIME_XML)//主路径
	//GET路由路径信息
	ws.Route(ws.GET("/{ID}").To(p.getOne)).
		Doc("Getting product id").
		Param(ws.PathParameter("id","identifier of the product").DataType("string"))
	//POST路由路径信息
	ws.Route(ws.POST("").To(p.postOne)).
		Doc("Updating product id").
		Param(ws.PathParameter("product","a product").DataType("main.product"))

	restful.Add(ws)//Webservice加入默认服务器
}
func main() {
	ProductResource{}.Register()
	log.Fatal(http.ListenAndServe(":8016",nil))
}
