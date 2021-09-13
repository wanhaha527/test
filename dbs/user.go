package dbs
//操作数据库表
import (

	"fmt"
)
var LastId int64
type User struct {
	id int
	name string
	age int
	gender string
	grade int
	address string
}
//查询数据

func QueryAll(){
	user:=new(User)
	stmt,_:=mysqldb.Prepare(`select * from user_test`)
	defer stmt.Close()
	rows,err:=stmt.Query()
	//fmt.Print(rows)
	if err!=nil {
		fmt.Println("select err")
	}

	Usr :=make([]User,0)
	for rows.Next() {
		rows.Scan(&user.id,&user.name,&user.age,&user.gender,&user.grade,&user.address)
		Usr =append(Usr,*user)

	}
	fmt.Println(Usr)


}
//插入数据

func InsertByLastId() {
	//user:=new(User)
	stmt,_:=mysqldb.Prepare(`insert into user_test (name ,age,gender,grade,address) values (?,?,?,?,?)`)
	defer stmt.Close()
	res,err:=stmt.Exec("小红",22,"男",90,"重庆")
	LastId,_ =res.LastInsertId()
	fmt.Printf("lastid=%d所在行",LastId)
	if err!=nil {
		fmt.Println("插入失败")
	}
	fmt.Println("插入成功")
}
func InsertById(id int) {
	//user:=new(User)
	stmt,_:=mysqldb.Prepare(`insert into user_test (id,name ,age,gender,grade,address) values (?,?,?,?,?,?)`)
	defer stmt.Close()
	res,err:=stmt.Exec(id,"小林",20,"男",90,"云南")
	LastId,_ =res.LastInsertId()
	fmt.Printf("id=%d所在行",id)
	if err!=nil {
		fmt.Println("插入失败")
	}
	fmt.Println("插入成功")
}

//删除数据

func DeleteById(Id int) {
	stmt,_:=mysqldb.Prepare(`delete from user_test where id=?`)
	defer stmt.Close()
	//const Id=19
	_,err:=stmt.Exec(Id)
	if err!=nil {
		fmt.Println("删除失败")
	}
	fmt.Printf("id=%d所在行删除成功",Id)
}
//更新数据

func UpdateById(Id int) {
	stmt,_:=mysqldb.Prepare(`update user_test set address=? where id=?`)
	defer stmt.Close()
	//const Id=19
	_,err:=stmt.Exec("绵阳",Id)
	if err!=nil {
		fmt.Println("更新失败")
	}
	fmt.Printf("id=%d所在行更新成功",Id)
}