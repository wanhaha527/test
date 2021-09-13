package main

import (
	"fmt"
	"time"
)

func main() {
	ch2 := make(chan int)
	done2 := make(chan struct{})
	//0,1,2
	for i := 0; i < 3; i++ {
		go func(idx int) {
			select {
			case ch2 <- (idx + 1) * 2://2,4,6
				fmt.Println(idx, "Send result")//2
			case <-done2://{}
				fmt.Println(idx, "Exiting")//0,1
			}
		}(i)
	}

	fmt.Println("Result: ", <-ch2)//6
	close(done2)
	time.Sleep(3 * time.Second)

	data := []int{1, 2, 3}
	for _, v := range data {//range迭代得到的值是元素值拷贝
		v *= 10        // data 中原有元素是不会被修改的
	}
	fmt.Println("data: ", data)    // data:  [1 2 3]

loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			//break    // 死循环，一直打印 breaking out...
			break loop
		}
	}
	fmt.Println("out...")

	data1 := []string{"one", "two", "three"}

	for _, v := range data1 {
		go func() {
			fmt.Println(v)
		}()
	}

	time.Sleep(3 * time.Second)

	data2 := [] string {"one", "two", "three"}
	for _, v := range data2 {
		//v := v
		fmt.Println(v)
	}
	time.Sleep(3 * time.Second)


	defer func() {
		fmt.Println("recovered: ", recover())//recover仅在defer中使用
	}()
	panic("not good")



}
