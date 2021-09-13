package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main()  {

	http.HandleFunc("/login1", login1)
	http.HandleFunc("/login2", login2)
	http.ListenAndServe(":8080", nil)
}

type Resp struct {
	Code	string `json:"code"`
	Msg		string `json:"msg"`
}

type  Auth struct {
	Username string `json:"username"`
	Pwd 	 string   `json:"password"`
}

//post接口接收json数据
func login1(writer http.ResponseWriter,  request *http.Request)  {
	var auth Auth
	if err := json.NewDecoder(request.Body).Decode(&auth); err != nil {
		request.Body.Close()
		log.Fatal(err)
	}
	var result  Resp
	if auth.Username == "admin" && auth.Pwd == "123456" {
		result.Code = "200"
		result.Msg = "登录成功"
	} else {
		result.Code = "401"
		result.Msg = "账户名或密码错误"
	}
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Fatal(err)
	}
}

//接收x-www-form-urlencoded类型的post请求或者普通get请求
func login2(writer http.ResponseWriter,  request *http.Request)  {
	request.ParseForm()
	username, uError :=  request.Form["username"]
	pwd, pError :=  request.Form["password"]

	var result  Resp
	if !uError || !pError {
		result.Code = "401"
		result.Msg = "登录失败"
	} else if username[0] == "admin" && pwd[0] == "123456" {
		result.Code = "200"
		result.Msg = "登录成功"
	} else {
		result.Code = "203"
		result.Msg = "账户名或密码错误"
	}
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Fatal(err)
	}
}



/*
var id int
var arr1[] int
func main() {
 arr:=[] int{1,2,3,4,5,6,7,8}

 fmt.Print(arr[:2])
 fmt.Print(arr[3:])
	arr1=append(arr[:2],arr[3:]...)//...将切片打散成一个个数据附加

 fmt.Print(arr1)
}
*/