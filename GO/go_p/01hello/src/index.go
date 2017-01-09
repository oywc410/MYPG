package main // 代码包声明语句。

// 代码包导入语句。
import (
	"fmt" // 导入代码包fmt。
)

// main函数。
func main() {

	var (
		num1 int = 1
		num2 int = 2
		num3 int = 3
		//其中计算架构相关的整数类型有两个，即：有符号的整数类型int和无符号的整数类型uint。
		// int8/16/32/64
		num uint64 = 2343
	)

	// 短变量声明语句，由变量名size、特殊标记:=，以及值（需要你来填写）组成。
	size := 64

	// 打印函数调用语句。用于打印输出信息。
	fmt.Println("Go语言编程实战")

	// 打印函数调用语句。用于打印上述三个变量的值。
	fmt.Println(num1, num2, num3)

	fmt.Printf("类型为 uint64 的整数 %d 需占用的存储空间为 %d 个字节。\n", num, size)

	// 声明一个整数类型变量并赋值
	num1 = -0x1000
	/*
		最显而易见的是10进制表示法。如果我们要在声明一个名为num1的int类型变量时给它赋予12这个值，那么这样书写即可：
		var num1 int = 12
	    这是最容易被我们理解的方式了。不过，如果我们要分别以8进制和16进制为变量num1赋值，那么需要：
		num1 = 014 // 用“0”作为前缀以表明这是8进制表示法。
		或：
		num1 = 0xC // 用“0x”作为前缀以表明这是16进制表示法。
	*/

	// 这里用到了字符串格式化函数。其中，%X用于以16进制显示整数类型值，%d用于以10进制显示整数类型值。
	fmt.Printf("16进制数 %X 表示的是 %d。\n", num1, num1)

	// 浮点数类型有两个，即float32和float64
	// 可以在变量声明并赋值的语句中，省略变量的类型部分。
	// 不过别担心，Go语言可以推导出该变量的类型。
	var num5 = 5.89E-4

	// 这里用到了字符串格式化函数。其中，%E用于以带指数部分的表示法显示浮点数类型值，%f用于以通常的方法显示浮点数类型值。
	fmt.Printf("浮点数 %E 表示的是 %f。\n", num5, num5)

	//复数类型
	/*
	 复数类型同样有两个，即complex64和complex128。存储这两个类型的值的空间分别需要8个字节和16个字节。实际上，complex64类型的值会由两个float32类型的值分别表示复数的实数部分和虚数部分。而complex128类型的值会由两个float64类型的值分别表示复数的实数部分和虚数部分。
	*/
	var num6 = 3.7E+1 + 5.98E-2i

	// 这里用到了字符串格式化函数。其中，%E用于以带指数部分的表示法显示浮点数类型值，%f用于以通常的方法显示浮点数类型值。
	fmt.Printf("浮点数 %E 表示的是 %f。\n", num6, num6)

	//byte与rune
	// byte与rune类型有一个共性，即：它们都属于别名类型。byte是uint8的别名类型，而rune则是int32的别名类型。
	// 声明一个rune类型变量并赋值
	var char1 rune = '赞'

	// 这里用到了字符串格式化函数。其中，%c用于显示rune类型值代表的字符。
	fmt.Printf("字符 '%c' 的Unicode代码点是 %s。\n", char1, "U+8D5E")

	// 声明一个string类型变量并赋值
	var str1 string = "\\\""

	// 这里用到了字符串格式化函数。其中，%q用于显示字符串值的表象值并用双引号包裹。
	fmt.Printf("用解释型字符串表示法表示的 %q 所代表的是 %s。\n", str1, `\"`)

	var numbers2 [5]int
	numbers2[0] = 2
	numbers2[3] = numbers2[0] - 3
	numbers2[1] = numbers2[2] + 5
	numbers2[4] = len(numbers2)
	sum := 11
	// “==”用于两个值的相等性判断
	fmt.Printf("%v\n", (sum == numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]))

	//截取数组
	//cap数组容量
	var numbers3 = []int{1, 2, 3, 4, 5}
	slice3 := numbers3[2:5]
	length := 3
	capacity := 3
	fmt.Printf("%v, %v, %d\n", (length == len(slice3)), (capacity == cap(slice3)), cap(slice3))

	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice5 := numbers4[4:6:8]
	fmt.Println(slice5, len(slice5), cap(slice5))
	length2 := 2
	capacity2 := 4
	fmt.Printf("%v, %v\n", length2 == len(slice5), capacity2 == cap(slice5))

	slice5 = slice5[:cap(slice5)]
	fmt.Println(slice5, len(slice5), cap(slice5))
	slice5 = append(slice5, 11, 12, 13)
	length = 7
	fmt.Printf("%v\n", length == len(slice5))

	slice6 := []int{0, 0, 0}
	copy(slice5, slice6)
	e2 := 0
	e3 := 8
	e4 := 9
	fmt.Println(slice5)
	fmt.Printf("%v, %v, %v\n", e2 == slice5[2], e3 == slice5[3], e4 == slice5[4])

	//int 类型值时默认无值是为0  string类型时默认值为""
	mm2 := map[string]int{"golang": 42, "java": 1, "python": 8, "scala": 25, "erlang": 50}
	fmt.Printf("%d, %d, %d \n", mm2["scala"], mm2["erlang"], mm2["aaaa"])

	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var mySlice []int = myArray[:5]
	fmt.Println(mySlice)

	myslice1 := make([]int, 5)
	myslice2 := make([]int, 5, 10) //初始个数5，预留10个元素的存储空间
	myslice3 := []int{1, 2, 3, 4, 5}

	fmt.Println(cap(myslice2))
	fmt.Println(myslice1, myslice2, myslice3)

	//append动态增加 append(array,array) append(array,int...)
	//cop 复制相同KEY
	//切片:
	//默认开始位置0，ar[:n]等价于ar[0:n]
	//第二个序列默认是数组长度 ar[n:] 等价于 ar[n:len(ar)]
	//从一个数组直接获取slice，可以是ar[:]

	//并发安全类型!!!:chan
	ch2 := make(chan string, 1)
	// 下面就是传说中的通过启用一个Goroutine来并发的执行代码块的方法。
	// 关键字 go 后跟的就是需要被并发执行的代码块，它由一个匿名函数代表。
	// 对于 go 关键字以及函数编写方法，我们后面再做专门介绍。
	// 在这里，我们只要知道在花括号中的就是将要被并发执行的代码就可以了。
	go func() {
		ch2 <- "已到达!" //在ch2中写入数据
	}()
	var value string = "数据"
	value = value + <-ch2 //从ch2中读出数据 (数据输出后 长度度为0)
	fmt.Println(value)

	ch2 <- "1"
	value2, ok := <-ch2
	fmt.Println(value2, ok)

	close(ch2) //关闭

}
