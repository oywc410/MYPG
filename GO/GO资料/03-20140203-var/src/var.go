package src

import (

)

//变量类型

var v1 int
var v2 string
var v3 [10]int					//数组
var v4 []int					//数组切片
var v5 struct {
	f int
}
var v6 *int						//指针
var v7 map[string]int			//map, key为string类型
var v8 func(a int) int

var (							//重复声明变量
	v9 int
	v10 string
)

//初始化变量
var v11 int = 10
var v12 = 10
v13 := 10						//自动识别类型


//变量赋值
var v14 int
v14 = 123


//多重赋值(交互i和j变量)
i, j = j, i

//匿名变量 获取多个返回值
func GetName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi Maruko"
}

firstName, lastName, nickName := GetName()
_, _, nickName := GetName() //只获取nickName



//-------常量----------------
//常量定义
const Pi float64 = 3.4134365634
const zero = 0.0				//无类型赋值
const (
	size int64 = 1024
	eof = -1
)

const u, v float32 = 0, 3		//u = 0.0   v = 3.0		常量的多重赋值
const a, b, c = 3, 4, "foo"		//a = 3 b = 4 c = "foo"

const mask = 1 << 3

//go 的预定义变量
//iota 自增变量
const (							//iota被重设为0
	c0 = iota					//c0 = 0
	c1 = iota					//c1 = 1
	c2 = iota					//c2 = 2
)

const (							//iota被重设为0
	a = 1 << iota				//a = 1
	b = 1 << iota				//b = 2
	c = 1 << iota				//c = 4
)

const x = iota					//x = 0
const y = iota					//y = 0

const (
	c0 = iota					//c0 = 0
	c1							//c1 = 1
	c2							//c2 = 2
)

//-----枚举------------
//以大写字母开头的常量在包外可见
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays // 这个常量没有导出
)

//-----变量类型-----------
布尔类型 bool
整型 int8 byte int16 int uint uintptr
浮点类型 float32 float64
复数类型 complex64 complex128
字符串 string
字符类型 rune
错误类型 error
符合类型 
	pointer 	指针
	array		数组
	slice		切片
	map			字典
	chan		通道
	struct		结构体
	interface	接口
	
	
//布尔类型
var v1 bool
v1 = true
v2 := (1 == 2)


//变量间的比较
var i int32
var j int64

if i == j {//编译错误
}

if i == 1 || j == 2 {//编译通过
}

//浮点数之间的比较
//p为用户自定义的比较精度,比如0.000001
func IsEqual(f1, f2, p float64) bool {
	return math.Fdim(f1, f2) < p
}

//字符串相关
var str string
str = "Hello world"
l = len(str)
ch := str[0]			//ch = H
fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)

//字符串操作
x + y 字符串连接
len(s) 字符串长度
s[i] 取字符

//遍历字符串   p44
//1 字节数组方式
str := "Hello,世界"
n := len(str)
for i := 0; i < n; i++ {
	ch := str[i]
	fmt.Println(i, ch)
}

//2 Unicode字符方式
str := "Hello,世界"
for i, ch := range str {
	fmt.Println(i, ch)
}



//数组
[32]byte					//长度为32的数组,每个元素为一个字节
[2*N] struct {x ,y int32}	//复杂类型的数组
[1000]*float64				//指针数组
[3][5] int					//二维数组
[2][2][2]float32			//等同于[2]([2]([2]float32))

arrLenght := len(arr)		//数组长度

for i, v := range arr {
	fmt.Println("Array element[", i, "]=", v)
}

