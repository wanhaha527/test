//操作数据库的方法

package db

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"io"
	"net/http"
	"strconv"
)


// GET http://localhost:8080/users
//查询所有用户

func  FindAllUsers(request *restful.Request, response *restful.Response) {

	stmt,_:= Mysqldb.Prepare(`select * from user_test`)
	defer stmt.Close()
	rows,err:=stmt.Query()
	if err!=nil {
		fmt.Println("插入错误")
	}
	user:= User{}
	Usr :=make([]User,0)
	for rows.Next() {

		rows.Scan(&user.Id,&user.Name,&user.Age,&user.Gender,&user.Grade,&user.Address)
		Usr =append(Usr,user)
	}
	fmt.Println(Usr)
	response.WriteEntity(Usr)
	/*
	   list := []User{}
	   	for _, each := range u.users {
	   		list = append(list, each)
	   	}
	   	response.WriteEntity(list)

	*/
}

// GET http://localhost:8080/users/lastid
//根据id查询用户

func  FindUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	ID, _:= strconv.Atoi(id)//int Id
	user:=User{}
	stmt,_:= Mysqldb.Prepare(`select id,name,age,gender,grade,address from user_test where id=?`)
	defer stmt.Close()
	_,err:=stmt.Exec(ID)
	err1:= stmt.QueryRow(ID).Scan(&user.Id, &user.Name, &user.Age, &user.Gender, &user.Grade, &user.Address)
	if err!=nil {
		fmt.Println("执行查询错误")
	}else {
		fmt.Printf("%d,%s,%d,%s,%d,%s\n",user.Id,user.Name,user.Age,user.Gender,user.Grade,user.Address)
	}
	if err1!=nil{
		fmt.Println("查询扫描错误")
		io.WriteString(response.ResponseWriter,"用户不能找到")
	}else {
		response.WriteEntity(user)
	}

}

// PUT http://localhost:8080/users/1
//根据id更新用户

func  UpdateUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	ID, _:= strconv.Atoi(id)//int id

	stmt,_:=Mysqldb.Prepare(`update user_test set address=? where id=?`)
	defer stmt.Close()
	_,err:=stmt.Exec("根据id更新的地点",ID)
	if err!=nil {
		fmt.Println("更新失败")
	}else {
		fmt.Printf("id=%d所在行更新成功\n",ID)
		io.WriteString(response.ResponseWriter,"更新完成")
	}

	/*usr := new(User)
	err := request.ReadEntity(&usr)
	if err == nil {
		u.users[usr.id] =*usr
		response.WriteEntity(usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}*/
}

// POST http://localhost:8080/users/add
//末尾行创建用户 (u *UserResource)

func  CreateUser(request *restful.Request, response *restful.Response) {
	//name,age,gender,grade,address
	stmt,_:=Mysqldb.Prepare(`insert into user_test (name,age,gender,grade,address) values (?,?,?,?,?)`)
	defer stmt.Close()
	res,err:=stmt.Exec("增加姓名",0,"男",0,"新增地点")
	var LastId,_ =res.LastInsertId()
	fmt.Printf("lastid=%d所在行",LastId)
	if err!=nil {
		fmt.Println("插入失败")
		io.WriteString(response.ResponseWriter,"末尾行添加用户失败")
	}else {
		fmt.Println("插入成功")
		io.WriteString(response.ResponseWriter,"末尾行添加用户成功")
	}
	//usr := User{name: request.PathParameter("user-name")}//name:user-name
	//usr:=new(User)
	//uid:=request.PathParameter("user-id")
	/*
			usr:=User{id:request.PathParameter("user-id")}
			err:=request.ReadEntity(&usr.id)
			if err==nil {
				u.users[usr.id]=usr
				response.WriteEntity(usr)
			}else {
				response.AddHeader("Content-Type","text/plain")
				response.WriteErrorString(http.StatusInternalServerError,err.Error())
			}

		if err!=nil {
				fmt.Println("插入失败")
			}

			err1 := request.ReadEntity(&usr)
			if err1 == nil {
				u.users[usr.name] = *usr//User{name:user-name}
				response.WriteEntity(usr)
			} else {
				response.WriteError(http.StatusInternalServerError, err1)
			}
	*/


}
//POST http://localhost:8080/users/add/{user-id}
//根据id创建用户

func  CreateUserById(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	ID, _:= strconv.Atoi(id)//int Id
	stmt, _ := Mysqldb.Prepare(`insert into user_test (id,name,age,gender,grade,address) values (?,?,?,?,?,?)`)
	defer stmt.Close()
	_, err := stmt.Exec(ID,"根据id增加的姓名", 10, "男", 99, "默认地点")
	if err != nil {
		fmt.Println("根据id新增行失败" )
	}
	fmt.Printf("根据id=%d新增所在行成功\n", ID)


	usr := User{Id: ID}
	err1 := request.ReadEntity(&usr)
	if err1 == nil {
		//u.users[usr.Id] = usr
		response.WriteHeaderAndEntity(http.StatusCreated, usr)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}

	/*
		usr:=User{id:Id}
		err1:=request.ReadEntity(&usr.id)
		if err1==nil {
				//u.users[usr.id]=usr
				response.WriteEntity(usr)
			}else {
				//response.AddHeader("Content-Type","text/plain")
				response.WriteErrorString(http.StatusInternalServerError,err.Error())
			}
	*/
}

// DELETE http://localhost:8080/users/1
//根据id删除用户

func  RemoveUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	//delete(u.users, id)//map[id]User
	stmt,_:=Mysqldb.Prepare(`delete from user_test where id=?`)
	defer stmt.Close()
	//const Id=19
	_,err:=stmt.Exec(id)
	if err!=nil {
		fmt.Println("删除失败")
		io.WriteString(response.ResponseWriter,"删除失败")
	}else{
		fmt.Printf("id=%s所在行删除成功\n",id)
		io.WriteString(response.ResponseWriter,"根据id删除成功")
	}

}
