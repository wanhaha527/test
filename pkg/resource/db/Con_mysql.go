//连接数据库

package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Mysqldb *sql.DB
var err error
const (
	root="root"
	password="wang5272357"
	localhost="172.25.208.1"//"mysqlServiceHost"//"172.17.0.1"////"172.17.0.1""host.docker.internal"
	port="3306"
	database="demo"
	charset="utf8mb4"
)

func init() {


	db:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", root, password, localhost, port, database, charset)
	Mysqldb, err =sql.Open("mysql",db)
	if err !=nil {
		fmt.Println("Open err")
	}
	if mysqlErr:= Mysqldb.Ping();mysqlErr!=nil {
		panic("数据库连接失败"+mysqlErr.Error())
	}

}
