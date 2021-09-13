package dbs
//初始化连接数据库
import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"

)
var mysqldb *sql.DB
var err error
const (
	root="root"
	password="wang5272357"
	localhost="localhost"
	port="3306"
	database="demo"
	charset="utf8mb4"
)

func init() {

	db:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",root,password,localhost,port,database,charset)

	mysqldb,err=sql.Open("mysql",db)
	if err!=nil {
		fmt.Println("Open err")
	}
	if mysqlErr:=mysqldb.Ping();mysqlErr!=nil {
		panic("数据库连接失败"+mysqlErr.Error())
	}


}