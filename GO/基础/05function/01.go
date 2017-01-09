package main

import "fmt"

// 定义无参数也无返回值的函数
func NoParaNoReturn() {
	fmt.Println("I am a function with no parameter and return nothing.")
}

// 定义有参数但是无返回值的函数
func WithParasButNoReturn(intPara1, intPara2 int, strPara string) {
	fmt.Println("My parameters are: ", intPara1, intPara2, strPara)
}

// 定义有参数也有返回值的函数
func WithpParasAndReturn(intPara int) (int, int) {
	return intPara, intPara * 2
}

//定义有参数也有含命名返回值的函数，return 的时候会默认返回命名返回值 doubledInput
func WithpParasAndNamedReturn(intPara int) (doubledInput int) {
	fmt.Println("Initial value of doubledInput of int type is: ", doubledInput)
	doubledInput = intPara * 2
	return
}

// 定义有参数也有含命名返回值的函数，如果内部声明了和命名返回值同名的局部变量，需要显式返回命名返回值
func WithpNamedReturnAndInnerVar(intPara int) (output int) {
	{
		var output = 100
		output = intPara * 2
		return output
	}
}

// 定义有 可变参数 的函数
func WithVariabicPara(prefix string, paras ...int) {
	tempInt := 0
	for _, para := range paras {
		tempInt += para
	}
	fmt.Printf("%s, The sum of paras is: %d\n", prefix, tempInt)
}

// 定义含有闭包或者说匿名函数的函数
func WithClosure(intPara int) func() {
	returnFunc := func() {
		fmt.Println("Input is: ", intPara)
	}
	return returnFunc
}

// 定义含有 延时代码 的函数
func DelayCode() {
	defer fmt.Println("Defer: I will be executed at last")
	defer fmt.Println("Defer: I will be executed at first")

	fmt.Println("I will be executed before defer")
}

// 定义含有回调函数的函数
func UseCallback(myFunc func()) {
	myFunc()
}

// 定义递归函数
func UseRecursion(num int) int {
	if num <= 0 {
		return 0
	} else if num == 1 || num == 2 {
		return 1
	} else {
		return UseRecursion(num-1) + UseRecursion(num-2)
	}
}

// 定义结构体
type User struct {
	name string
}

// 定义属于结构体 User 的方法
func (user *User) NameChangedTo(new string) {
	user.name = new
}

// 定义属于结构体 User 的另一个方法
func (user *User) GetName() string {
	return user.name
}

func main() {
	// 调用无参数也无返回值的函数
	NoParaNoReturn()

	// 调用有参数但是无返回值的函数
	WithParasButNoReturn(1, 2, "OK")

	// 调用有参数也有返回值的函数
	input, output := WithpParasAndReturn(10)
	fmt.Printf("The double of %d is: %d\n", input, output)

	// 调用有参数也有含命名返回值的函数
	fmt.Println("The double of 10 is: ", WithpParasAndNamedReturn(10))

	// 调用有参数也有含命名返回值以及同名局部变量的函数
	fmt.Println("The double of 10 is: ", WithpNamedReturnAndInnerVar(10))

	// 调用 可变参数 的函数
	WithVariabicPara("Test", 1, 2, 3)

	// 调用递归函数
	fmt.Printf("The %d num of Fabinaci is: %d\n", 5, UseRecursion(5))

	// 调用含有闭包的函数
	WithClosure(100)()

	// 调用含有 DelayCode 的函数
	DelayCode()

	// 调用含有回调函数的函数
	UseCallback(DelayCode)

	// 调用 struct 的方法
	var user User
	fmt.Println("user is: ", user)
	user.NameChangedTo("new")
	fmt.Println("user.name is: ", user.GetName())

}
