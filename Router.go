package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	//StrictSlash定义尾斜杠行为。初始值为false。
	//如果为true，则如果路由路径为“/path/”，则访问“/path”将执行重定向到前者，反之亦然
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		//处理函数
		handler = route.HandlerFunc
		//打印日志
		handler = logger(handler, route.Name)
	//路由器定义
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}