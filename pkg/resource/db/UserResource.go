// 定义资源结构

package db

type UserResource struct {
	// 数据访问对象(data access object)
	users map[int]User
}
