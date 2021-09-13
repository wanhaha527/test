package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitemty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City     string `json:"city,omitempty"`
	Province string `json:"province,omitempty"`}

var people []Person

// *******************************************************************>>
// Get
// 获取所有person

func GetPeople(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
	log.Print("获取所有person")
}

// 根据id获取person

func GetPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {//获取获得的参数id
			json.NewEncoder(w).Encode(item)
			log.Printf("根据id%s获取了一条数据",item.ID)
			return
		}
	}
	json.NewEncoder(w).Encode(people)

}

// <<*******************************************************************

// *******************************************************************>>
// Post
// 通过post操作向服务器添加数据

func PostPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	fmt.Print("向服务器添加了一条数据\n")
	json.NewEncoder(w).Encode(people)

}

// <<*******************************************************************

// *******************************************************************>>
// Delete
// 根据id进行删除操作

func DeletePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			fmt.Print("删除了一条数据\n")
			break
		}
	}
	json.NewEncoder(w).Encode(people)

}

// <<*******************************************************************

func main() {
	people = append(people, Person{ID: "1", Firstname: "xi", Lastname: "dada", Address: &Address{City: "Shenyang", Province: "Liaoning"}})
	people = append(people, Person{ID: "2", Firstname: "li", Lastname: "xiansheng", Address: &Address{City: "Changchun", Province: "Jinlin"}})

	// Get handle function:
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")

	// Post handle function
	router.HandleFunc("/people/{id}", PostPerson).Methods("POST")

	// Delete handle function:
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	// 启动 API端口9899
	log.Fatal(http.ListenAndServe(":9899", router))
}
