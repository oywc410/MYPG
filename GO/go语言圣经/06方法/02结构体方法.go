package main
import "fmt"


type P struct {

}

func (p P) test() {
	fmt.Println("test1")
}

func (p *P) test2() {
	fmt.Println("test2")
}

type P2 struct  {
	P
}

func main() {
	//结构体
	p := P{}
	p.test2()

	//结构体指针
	p2 := P2{}
	p2.test2()

	//获取结构体方法与使用
	testfunc := P.test
	testfunc(p)
}
