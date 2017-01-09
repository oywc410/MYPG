package main

import "fmt"

type Pet interface {
    Name() string
    Age() uint8
}

type Dog struct{
    DogName string
    DogAge uint8
}

func (dog Dog)Name ()(string){
    return dog.DogName
}
func (dog Dog)Age ()(uint8){
    return dog.DogAge
}

type abc struct{
    v int
}

func (a abc)aaaa(){ //传入的是值，而不是引用
    a.v=1
    fmt.Printf("1:%d\n",a.v)
}
 
func (a *abc)bbbb(){ //传入的是引用，而不是值
    fmt.Printf("2:%d\n",a.v)
    a.v=2
    fmt.Printf("3:%d\n",a.v)
}
 
func (a *abc)cccc(){ //传入的是引用，而不是值
    fmt.Printf("4:%d\n",a.v)
}

//运算符就是简单的 & 和 * 一个取地址、一个解析地址。

type S map[string][]string
 
func Summary(paramstring string)(s *S){
    s=&S{
        "name":[]string{paramstring},
        "profession":[]string{"Javaprogrammer","ProjectManager"},
        "interest(lang)":[]string{"Clojure","Python","Go"},
        "focus(project)":[]string{"UE","AgileMethodology","SoftwareEngineering"},
        "hobby(life)":[]string{"Basketball","Movies","Travel"},
    }
    return s
}

func main() {

	var i int // i 的类型是int型
    i=1 // i 的值为 1;
    var p *int // p 的类型是[int型的指针]
    p=&i         // p 的值为 [i的地址]
 
    fmt.Printf("i=%d;p=%d;*p=%d\n",i,p,*p)
 
    *p=2 // *p 的值为 [[i的地址]的指针] (其实就是i嘛),这行代码也就等价于 i = 2
    fmt.Printf("i=%d;p=%d;*p=%d\n",i,p,*p)
 
    i=3 // 验证想法
    fmt.Printf("i=%d;p=%d;*p=%d\n",i,p,*p)
	
	//-------------------------------
	
	aobj:=abc{}  //new(abc);
    aobj.aaaa()
    aobj.bbbb()
    aobj.cccc()
	
	
	//------------------------------
	
	s:=Summary("Harry")
    fmt.Printf("Summary(address):%v\r\n",s)
    fmt.Printf("Summary(content):%v\r\n",*s)
	
	//-----------------------------

	myDog := Dog{"Little D", 3}
    
	pet1, ok1 := interface{}(&myDog).(Pet)
    pet1.Name()
    fmt.Printf(myDog.DogName + "\n")
    
    pet2, ok2 := interface{}(myDog).(Pet)
    pet2.Name()
    fmt.Printf(myDog.DogName + "\n")
    
	fmt.Printf("%v, %v\n", ok1, ok2)
}