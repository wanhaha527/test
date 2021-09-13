//被go-restful使用，作为操作的对象，对应了API文档的definitions部分

package db

type User struct {
	Id   int `json:"id" description:"identifier of the user"`
	Name string `json:"name" description:"name of the user"`
	Age  int    `json:"age" description:"age of the user" default:"0"`
	Gender string  `json:"gender" description:"gender of the user"`
	Grade int      `json:"grade" description:"grade of the user"`
	Address string `json:"address" description:"address of the user"`
}



/*type User struct {
	Id   int `json:"id" description:"identifier of the user"`
	Name string `json:"name" description:"name of the user"`
	Age  int    `json:"age" description:"age of the user" default:"0"`
	Gender string  `json:"gender" description:"gender of the user"`
	Grade int      `json:"grade" description:"grade of the user"`
	Address string `json:"address" description:"address of the user"`

}
type UserResource struct {
	// (data access object)
	users map[int]User
	//User
}*/

