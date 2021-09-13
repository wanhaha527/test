package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Printf("请输入内容：")
	fmt.Scanf("%s", &str)
	if _,err:=strconv.Atoi(str);err!=nil {
		fmt.Println("错误，转换失败")
	}else{
		fmt.Println("是数字")
	}
	fmt.Printf("输入了：%s", str)
}

