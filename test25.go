package main

import (
	"encoding/json"
	"fmt"
)

import (
	"errors"
	"strconv"

	//"golang.org/x/text/encoding/korean"
	"log"
)

//利用接口实现多态
//序列化及反序列化，序列化简单类型，序列化结构化数据，编写自定义Marshaller
func main() {
	test(new(People))

	jsonStr := `{"number":2300567}`
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &result)//反序列化
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	foo := func() {
		fmt.Println("foo() here")
	}

	_, err = json.Marshal(foo)//不能反序列化函数
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("*********************************")
	//map序列化与反序列化
	dd:=`{
		"label":"root",
		"value":3,
		"left":{
			"value":1,
			"left":null,
			"right":{
				"value":2,
				"left":null,
				"right":null
			}

		},
		"right":{
			"value":4,
			"left":null,
			"right":null
		}
	}`
	var obj interface{}
	err=json.Unmarshal([]byte(dd),&obj)
	if err!=nil{
		fmt.Println("unmarshal err")
	}else{
		fmt.Println("------\n",obj)
	}
	//var data
	data,err1:=json.Marshal(obj)
	if err1!=nil{
		fmt.Println("marshal err")
	}else{
		fmt.Println("------\n",string(data))
	}
	fmt.Println("***************************")
	dump(obj)
	fmt.Println("***************************")
	var tr tree
	err=json.Unmarshal([]byte(dd),&tr)
	if err!=nil{
		fmt.Printf("--connot deseriselize tree %v\n",err)
	}else{
		tr.Dump("")
	}
	fmt.Println("***************************")
	//自定义编组器
	//序列化
	m:=IntStringMap{4:"four",5:"five",6:"six"}
	data0,err0:=m.MarshalJson()
	if err!=nil {
		fmt.Println(err0)
	}else{
		fmt.Println("IntStringMap to Json",string(data0))
	}
	//反序列化
	m=IntStringMap{}
	JsonString:=[]byte("{\"1\":\"one\",\"2\":\"two\"}")
	m.UnmarshalJson(JsonString)
	fmt.Printf("IntStringMap from Json: %v\n",m)
	fmt.Println("m[1]",m[1],"m[2]",m[2])

	fmt.Println("test git")

}
// --- interface
type Animal interface {//接口
	say(words string)//方法
}

func test(s Animal) {
	s.say("animal")
}

type People struct {
}
func (p *People) say(word string) {
	log.Printf("你好 %v\n", word)
}



// 多态
/*
func TestSayHello(t *testing.T) {
	SayHello(new(People))
	SayHello(new(Dog))
	SayHello(new(Cat))
}*/
//遍历接口的通用映射，您需要使用类型断言。 例如
func dump(obj interface{}) error {
	if obj==nil{
		fmt.Println("nil")
		return nil
	}
	switch obj.(type) {
		case bool:
			fmt.Println(obj.(bool))
		case int:
			fmt.Println(obj.(int))
		case float64:
			fmt.Println(obj.(float64))
		case string:
			fmt.Println(obj.(string))
		case map[string]interface{}:
			for k,v:=range obj.(map[string]interface{}){
				fmt.Printf("%s",k)
				err:=dump(v)
				if err!=nil{
					return err
				}
			}
		default:
			return errors.New(fmt.Sprintf("unsupported type:%v",obj))
	}
	return nil
}

//序列化结构化数据
//“树”字段是私有的。 JSON序列化仅适用于公共字段。 因此，我们可以将struct字段公开
type tree struct {
	Tag string `json:"label"`
	Value int
	Left *tree
	Right *tree
}

func (t *tree) Dump(indent string) {
	fmt.Println(indent+"value:",t.Value)
	if t.Tag!="" {
		fmt.Println(indent+"tag:",t.Tag)
	}

	fmt.Println(indent+"left:")
	if t.Left==nil{
		fmt.Println("nil")
	}else{
		//fmt.Println()
		t.Left.Dump(indent+" ")
	}

	fmt.Println(indent+"right:")
	if t.Right==nil{
		fmt.Println("nil")
	}else{
		//fmt.Println()
		t.Right.Dump(indent+" ")
	}
}

//编写自定义编组器,关于拼写的注意事项：在Go中，约定是通过在方法名称后附加“ er”后缀来用单个方法命名接口
type Marshaler interface {
	MarshalJson()([]byte,error)
}
type Unmarshaler interface {
	UnmarshalJson([]byte)(error)
}
type IntStringMap map[int]string


func (m *IntStringMap) MarshalJson() ([]byte, error) {
	ss := map[string]string{}
	for k, v := range *m{
		i := strconv.Itoa(k)
		ss[i] = v
	}
	return json.Marshal(ss)
}

//反序列化
func (m *IntStringMap)UnmarshalJson(data []byte) error  {
	ss:=map[string]string{}
	err1:=json.Unmarshal(data,&ss)
	if err1!=nil{
		return err1
	}
	for k,v:=range ss{
		i,err2:=strconv.Atoi(k)
		if err2!=nil {
			return err2
		}
		(*m)[i]=v
	}
	return nil
}
