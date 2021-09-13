package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)

type DbWorker struct {
	Dsn      string
	Db       *sql.DB
	UserInfo userTB
}
type userTB struct {
	Id   int `demo:"id"`
	name sql.NullString
	age  sql.NullInt64
}

func main() {
	var err error
	dbw := DbWorker{
		Dsn: "root:wang5272357@tcp(localhost:3306)/demo?charset=utf8mb4",
	}
	dbw.Db, err = sql.Open("mysql", dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer dbw.Db.Close()

	dbw.insertData()
	dbw.queryData()
	//dbw.delete()
	//dbw.RawQueryField()

}
//插入数据id,name,age
func (dbw *DbWorker) insertData() {

	stmt, _ := dbw.Db.Prepare(`insert into  user(Id,name,age) VALUES (?,?,?)`)
	defer stmt.Close()

	ret, err := stmt.Exec(11,"ww", 23)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", LastInsertId)//使用自增主键才有返回值，否则为0
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:", RowsAffected)
		//fmt.Println(RowsAffected)
	}
}

func (dbw *DbWorker) QueryDataPre() {
	dbw.UserInfo =userTB{}
}
func (dbw *DbWorker) queryData() {


	stmt, _ := dbw.Db.Prepare(`SELECT * From user where age >= ? AND age < ?`)//注意字段顺序
	defer stmt.Close()

	dbw.QueryDataPre()
	//dbs:=new(userTB)
	rows, err := stmt.Query(50, 60)
	fmt.Println(rows)

	defer rows.Close()

	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	for rows.Next() {
		rows.Scan(&dbw.UserInfo.name, &dbw.UserInfo.age, &dbw.UserInfo.Id)
		//fmt.Println(dbw.UserInfo)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		if !dbw.UserInfo.name.Valid {
			dbw.UserInfo.name.String = ""
		}
		if !dbw.UserInfo.age.Valid {
			dbw.UserInfo.age.Int64 = 0
		}
		fmt.Println(" name: ",dbw.UserInfo.name.String, " age: ", int(dbw.UserInfo.age.Int64),"id: ",dbw.UserInfo.Id)
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}
func (dbw *DbWorker)delete(){
	_,err:=dbw.Db.Exec(`delete From user where age=56`)

	if err!=nil{
		fmt.Println("删除失败",err)

		return

	}
	fmt.Println("删除成功")

}
func (dbw *DbWorker) RawQueryField() {
	stmt, _ := dbw.Db.Prepare(`SELECT * From user where age >= ? AND age < ?`)
	defer stmt.Close()
	us:=new(userTB)
	rows, _ := stmt.Query(55, 60)
	//rows, _ := dbw.Db.Query("select * from user where age >=55 and age <60")
	if rows == nil {
		fmt.Println("rows==nil")
		return
	}
	//id := 0
	//name := ""
	fmt.Println(rows)
	for rows.Next() {
		rows.Scan(&us.name,&us.age,&dbw.UserInfo.Id)//分别使用us和dbw两种调用
		fmt.Println(us.name.String,us.age.Int64,dbw.UserInfo.Id)
	}
}
