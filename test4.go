package main
//装饰器模式
import "fmt"

type HandlerFunc func(a,b string)

func (f HandlerFunc) ServeTim(a,b string) {
	f(a, b)
}

type Handler interface {
	ServeTim(x, y string)
}

func Create(m,n string)  {
	fmt.Println(m,"Create",n)
}


func Delete(m,n string)  {
	fmt.Println(m,"Delete",n)
}

func main() {

	HandlerFunc(Create).ServeTim("333","444")

	HandlerFunc(Delete).ServeTim("333","444")
}