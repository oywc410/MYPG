package main

import "fmt"

// 定义空 `struct` 类型
type EmptyStruct struct{}

// 定义含有 `tag` 的 `struct` 类型，之后可以用 reflection 对有 `tag` 的属性进行访问。
type Job struct {
	title string "JD"
	rank  int
}

// 定义不含有 `tag`、其他属性和 Job `struct` 类型完全相同的另一个 `struct` 类型
type JobWithoutTag struct {
	title string
	rank  int
}

// 定义 `struct` 类型 Job 的方法，由于方法体中无需用到 Job 对象，那么对象名可以省略
func (Job) Help() {
	fmt.Println("Use \"j.What()\" to get job detail.")
}

// 定义 `struct` 类型 Job 的方法，由于方法体中用到 Job 对象，那么对象名不可以省略
func (job Job) What() {
	fmt.Printf("Job Detail:\n\ttitle: %s\n\trank: %d\n", job.title, job.rank)
}

// 定义指向 `struct` 类型 Job 的指针的方法
func (jobP *Job) SetRank(newRank int) {
	jobP.rank = newRank
	fmt.Println("Rank updated, new value is: ", jobP.rank)
}

// 定义 Employee 结构类型，其含有 `补位属性`、匿名 `struct` 类型属性、匿名 `struct` 类型而显式命名的属性
// 同一类型的属性可以放在一行
// 从效果看，Employee `继承` 了 Job，这和传统 `继承` 区别很大，注意上面提到的 `匿名属性` 的方法集规则
// 在其他 OOP 语言比如 C++ 中，我们不会做这样的抽象设计，因为两者没有什么所谓的 `继承关系`
type Employee struct {
	_                   int
	age                 int
	firstName, lastName string
	pack                struct {
		salary int
		stock  int
	}
	*Job
}

// 定义 `struct` 类型 Employee 的方法，由于其匿名属性 Job 也有同名方法，根据上面提到的规则
// Employee 的 Help 方法会 override 其匿名属性 Job 的同名方法。
func (Employee) Help() {
	fmt.Println("Use \"e.Job.What()\" to get job detail of current employee.")
}

// 定义 `struct` 类型，这个类型唯一的属性为空接口类型，而任何类型都实现了空接口
// 所以 OnlyInterface 可以实例化属性为任何类型的 `struct` 对象
type OnlyInterface struct {
	f interface{}
}

func main() {

	// 检查空 `struct` 类型
	varEmptyStruct := EmptyStruct{}
	fmt.Println("varEmptyStruct: ", varEmptyStruct)

	// 检查 `struct` 类型对象的默认值
	var varJob0 Job
	fmt.Println("varJob0 = ", varJob0)

	// 对 `struct` 类型对象进行顺序初始化
	varJob1 := Job{
		"CEO",
		99,
	}
	fmt.Println("varJob1 = ", varJob1)

	// 对 `struct` 类型对象按照属性名依次初始化，可以和属性定义的顺序不一致
	varJob2 := Job{
		rank:  99,
		title: "CEO",
	}
	fmt.Println("varJob2 = ", varJob2)

	// 使用属性名进行初始化可以选择性初始化个别属性，其他属性初始化为它们的默认值
	varJob3 := Job{
		title: "COO",
	}
	fmt.Println("varJob2 = ", varJob3)

	// Employee  `struct` 类型对象的默认值
	varEmployee0 := Employee{}
	fmt.Println("varEmployee = ", varEmployee0)

	// 初始化 Employee `struct` 类型对象，匿名属性需要在初始化其他属性后使用对象名显式赋值
	varEmployee1 := Employee{
		age:       50,
		firstName: "Jack",
		lastName:  "Ma",
	}
	varEmployee1.pack.salary = 100000000
	varEmployee1.pack.stock = 1000000
	varJob4 := Job{
		"Founder",
		100,
	}
	varEmployee1.Job = &varJob4
	fmt.Println("varEmployee = ", varEmployee1)
	fmt.Println("varEmployee.Job = ", *(varEmployee1.Job))

	// 注意 Employee 的 Help 方法 `override` 了其属性 Job 的 Help 方法
	varEmployee1.Help()
	// 调用 Job 的 Help 方法需要显式调用
	varEmployee1.Job.Help()
	// 使用 Employee 对象可以直接调用其属性 Job 的方法 What
	varEmployee1.What()
	// Employee 对象的属性 Job 的类型是 `*Job`，所以可以修改 Job 的内容
	// 如果其类型是 Job 本身，由于获取到的 `struct` 对象是其原始值的拷贝，修改不会生效
	varEmployee1.Job.SetRank(99)
	fmt.Println("After job rank change, varEmployee.Job = ", *(varEmployee1.Job))

	// `struct` 类型对象是值类型，可以进行比较运算
	varEmployee2 := Employee{
		age:       50,
		firstName: "Jack",
		lastName:  "Ma",
	}
	varEmployee2.pack.salary = 100000000
	varEmployee2.pack.stock = 1000000
	varEmployee2.Job = &varJob4
	if fmt.Println("Cmpare 2 struct object."); varEmployee2 == varEmployee1 {
		fmt.Println("varEmployee2 equals varEmployee1.")
	}

	// Tag 是 `struct` 类型的一部分，下面 varJob5 无法赋值给 varJob4
	// varJob5 := JobWithoutTag{
	// 	"Founder",
	// 	100,
	// }
	// varJob4 = varJob5

	// 属性类型为空接口的 `struct` 类型的使用
	varOnlyInterface := OnlyInterface{
		f: 100,
	}
	fmt.Println("OnlyInterface with int: ", varOnlyInterface)
	varOnlyInterface = OnlyInterface{
		f: "I like Golang.",
	}
	fmt.Println("OnlyInterface with string: ", varOnlyInterface)

}
