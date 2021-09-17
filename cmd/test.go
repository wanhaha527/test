package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"

	"testProject/pkg/resource/db"
)

type(
	UserResource db.UserResource
	User db.User
)


func main() {

	u := UserResource{}//map[int]User{}
	restful.DefaultContainer.Add(u.WebService())
	//config设置了生成API文档的一些配置
	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(), // 为哪个WebServices 生成 API文档
		APIPath:     "/test.json",//访问API文档的路径
		//enrichSwaggerObject 对应生成了swagger API 文档的info和 tags部分
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}//使用的基本信息（info和tags）
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))

	//打开http://localhost:3002/?url=http://localhost:8080/apidocs.json
	//http.Handle是这个程序和swagger ui在同一个服务器上，启动这个程序的时候同时启动swagger ui用的
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("/Users/wang/go/src/swagger-ui-master/dist"))))

	// cors 这个过滤器是为了让swagger ui能够访问到这个 API文档，没有的话，用swagger ui访问会报错
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		CookiesAllowed: true,
		Container:      restful.DefaultContainer}
	restful.DefaultContainer.Filter(cors.Filter)

	log.Printf("Get the API using http://localhost:8080/test.json")
	log.Printf("Open Swagger UI using http://127.0.0.1:3002/?url=http://localhost:8080/test.json")
	log.Fatal(http.ListenAndServe(":8080", nil))


}

func (u UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)	//Consumes,Produces对应swagger中的basePath和Consumes、Produces，在生成API过程中把他们分解到了具体的每一个path
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // 可以具体指定每一个route
	//tags 字符串数组，在ws.Route.Metadata中用到，规定了这个path的标签，
	//如果enrichSwaggerObject中定义的tags定义了这个tag,将从里面取对应的Description信息，如果没有定义，就当作一个全新的tag
	tags := []string{"users"}
	//ws.Route生成了API文档的path,ws.Route(ws.GET)对应了一个path的一个operation,
	//ws.Route.To生成API文档的时候，函数的名称会做为path.operation.operationId
	ws.Route(ws.GET("/").To(db.FindAllUsers).
		// docs
		Doc("get all users").//ws.Route.Doc对应生成了API文档的 path.operation.summary
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]User{}).//ws.Route.Writes生成path.operation.responses.default
		Returns(200, "OK", []User{}).//生成responses对应的状态码的返回。三个参数分别对应状态码，description，schema
		DefaultReturns("OK", []User{}))

	ws.Route(ws.GET("/{user-id}").To(db.FindUser).
		// docs
		Doc("get a user").
			//ws.Route.Param生成API文档的parameter
			Param(ws.PathParameter("user-id", "identifier of the user").
			DataType("integer").
			DefaultValue("1")).
			Metadata(restfulspec.KeyOpenAPITags, tags).
			Writes(User{}). // on the response
			Returns(200, "OK", User{}).
			Returns(404, "Not Found", nil).
			DefaultReturns("OK", User{}))

	ws.Route(ws.PUT("/{user-id}").To(db.UpdateUser).
		// docs
		Doc("update a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		//ws.Route.Reads也是生成API文档的parameter，区别是它使用了schema，引用了definitions中的对象
		Reads(User{})) // from the request

	ws.Route(ws.POST("/add/").To(db.CreateUser).
		// docs
		Doc("create a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(User{})) // from the request

	ws.Route(ws.POST("/add/{user-id}").To(db.CreateUserById).
		// docs
		Doc("create a user by id").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("integer")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(User{})) // from the request

	ws.Route(ws.DELETE("/{user-id}").To(db.RemoveUserById).
		// docs
		Doc("delete a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))

	ws.Route(ws.DELETE("/").To(db.RemoveUser).
		// docs
		Doc("delete tail user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]User{}).//ws.Route.Writes生成path.operation.responses.default
		Returns(200, "OK", []User{}).//生成responses对应的状态码的返回。三个参数分别对应状态码，description，schema
		DefaultReturns("Remove OK", []User{}))
	return ws
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "UserService",
			Description: "Resource for managing Users",
			Contact: &spec.ContactInfo{
				ContactInfoProps: spec.ContactInfoProps{
					Name:  "wangaiqin",
					Email: "aiqin.wang@changhong.com",
					URL:   "https://wangaiqin.com",
				},

			},
			License: &spec.License{
				LicenseProps: spec.LicenseProps{
					Name: "test",
					URL:  "https://test.org",
				},
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "users",
		Description: "Managing users"}}}
}



