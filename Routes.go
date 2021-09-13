package main
//定义路径信息
import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type R []Route//切片R

var routes = R{
	Route{Name:"Index",     Method:"GET",   Pattern:"/",                HandlerFunc:Index},
	Route{Name:"TodoIndex", Method:"GET",   Pattern:"/todos",           HandlerFunc:TodoIndex},
	Route{Name:"TodoShow",  Method:"GET",   Pattern:"/todos/{todoId}",  HandlerFunc:TodoShow},
}