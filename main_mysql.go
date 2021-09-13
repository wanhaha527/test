package main
//入口主函数
import "testmysql/dbs"

func main() {
dbs.InsertByLastId()
//dbs.InsertById(18)
dbs.QueryAll()
//dbs.DeleteById(25)
//dbs.UpdateById(24)
//dbs.DeleteById(17)
}
