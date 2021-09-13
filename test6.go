package main
// This example shows the minimal code needed to get a restful.WebService working.
import (
	"io"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)



func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8006", nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}