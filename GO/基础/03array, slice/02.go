package main

import "fmt"

func sliceTest(sli []int) {
	fmt.Printf("Address of sli is: %d\n", &sli[0])
}

func main() {

	// 声明数组
	v_IntArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("v_IntArray is: ", v_IntArray)

	// 对数组进行切片，方括号里面数字的意义为 `[low, high, capacity]`
	// 对应数组索引 `[starterIndex, endIndex + 1, lengthOfArray - starterIndex]`
	v_IntSlice := v_IntArray[1:4:9]
	fmt.Println("v_IntArray[1:4:9], v_IntSlice is: ", v_IntSlice)
	fmt.Println("v_IntArray[1:4:9], len(v_IntSlice) is: ", len(v_IntSlice))
	fmt.Println("v_IntArray[1:4:9], cap(v_IntSlice) is: ", cap(v_IntSlice))

	// 省略 `high` 和 `capacity`，省略 `high` 会从数组索引为 `low` 的元素一直到数组最后一个元素，切片的容量为 `lengthOfArray - starterIndex`
	v_IntSlice = v_IntArray[1:]
	v_IntSlice[0] = 100
	fmt.Println("v_IntArray is: ", v_IntArray)
	fmt.Println("v_IntArray[1:], v_IntSlice is: ", v_IntSlice)
	fmt.Println("v_IntArray[1:], len(v_IntSlice) is: ", len(v_IntSlice))
	fmt.Println("v_IntArray[1:], cap(v_IntSlice) is: ", cap(v_IntSlice))

	// 省略 `low` 和 `capacity`，省略 `low` 会从数组索引为 `0` 的元素一直到索引为 `high` 的元素
	// 切片的容量和数组长度相等
	v_IntSlice = v_IntArray[:1]
	fmt.Println("v_IntArray[:1], v_IntSlice is: ", v_IntSlice)
	fmt.Println("v_IntArray[:1], len(v_IntSlice) is: ", len(v_IntSlice))
	fmt.Println("v_IntArray[:1], cap(v_IntSlice) is: ", cap(v_IntSlice))

	// 全部省略，从数组索引为 `0` 的元素一直到数组最后一个元素，切片的容量和数组长度相等
	v_IntSlice = v_IntArray[:]
	fmt.Println("v_IntArray[:], v_IntSlice is: ", v_IntSlice)
	fmt.Println("v_IntArray[:], len(v_IntSlice) is: ", len(v_IntSlice))
	fmt.Println("v_IntArray[:], cap(v_IntSlice) is: ", cap(v_IntSlice))

	// 省略 `low`，省略 `low` 会从数组索引为 `0` 的元素一直到索引为 `high` 的元素
	// 切片对象的 `capacity` 为给定的数值，注意这个数值不能大于数组的长度
	v_IntSlice = v_IntArray[:4:4]
	fmt.Println("v_IntArray[:4:4], v_IntSlice is: ", v_IntSlice)
	fmt.Println("v_IntArray[:4:4], len(v_IntSlice) is: ", len(v_IntSlice))
	fmt.Println("v_IntArray[:4:4], cap(v_IntSlice) is: ", cap(v_IntSlice))

	// 如果所有初始值已经确定，可以这样直接声明 `slice` 对象
	v_IntSlice = []int{1, 2, 5: 6}
	v_IntSlice[0] = 100
	fmt.Println("v_IntArray is: ", v_IntArray)
	fmt.Println("[]int{1, 2, 5: 6}, v_IntSlice is: ", v_IntSlice)

	// 如果需要之后进行 `slice` 对象赋值，可以使用 `make` 内置函数声明 `slice` 对象
	v_IntSlice = make([]int, 5, 10)
	fmt.Println("make([]int, 5, 10), v_IntSlice is: ", v_IntSlice)
	fmt.Println("make([]int, 5, 10), len(v_IntSlice) is: ", len(v_IntSlice))
	fmt.Println("make([]int, 5, 10), cap(v_IntSlice) is: ", cap(v_IntSlice))

	// 使用 `make` 内置函数声明 `slice` 对象的时候可以省略 `capacity`
	// 其默认值和给定的 `slice` 的 `length` 相等。
	v_IntSlice = make([]int, 5)
	fmt.Println("make([]int, 5), v_IntSlice is: ", v_IntSlice)
	fmt.Println("make([]int, len(v_IntSlice) is: ", len(v_IntSlice))
	fmt.Println("make([]int, cap(v_IntSlice) is: ", cap(v_IntSlice))

	// 使用 `new` 内置函数也可以声明 `slice` 对象
	v_IntSlice = new([10]int)[:5]
	fmt.Println("new([10]int)[:5], v_IntSlice is: ", v_IntSlice)
	fmt.Println("new([10]int)[:5], len(v_IntSlice) is: ", len(v_IntSlice))
	fmt.Println("new([10]int)[:5], cap(v_IntSlice) is: ", cap(v_IntSlice))

	// 可以对切片进行再切片
	v_IntSlice = make([]int, 5)
	v_AnotherIntSlice := v_IntSlice[3:]
	v_AnotherIntSlice[0] = 100
	fmt.Println("v_AnotherIntSlice = v_IntSlice[3:], len(v_AnotherIntSlice) is: ", len(v_AnotherIntSlice))
	fmt.Println("v_AnotherIntSlice = v_IntSlice[3:], cap(v_AnotherIntSlice) is: ", cap(v_AnotherIntSlice))
	fmt.Println("v_AnotherIntSlice = v_IntSlice[3:], v_IntSlice is: ", v_IntSlice)
	fmt.Println("v_AnotherIntSlice = v_IntSlice[3:], v_AnotherIntSlice is: ", v_AnotherIntSlice)

	// 使用 `append` 内置函数可以对 `slice` 对象在其容量范围内进行添加，新添加的元素索引为 `len(切片对象)`，操作返回新的切片对象
	v_IntSlice = make([]int, 1, 2)
	v_AnotherIntSlice = append(v_IntSlice, 100)
	fmt.Println("v_AnotherIntSlice = append(v_IntSlice, 100), v_IntSlice is: ", v_IntSlice)
	fmt.Println("v_AnotherIntSlice = append(v_IntSlice, 100), v_AnotherIntSlice is: ", v_AnotherIntSlice)
	fmt.Printf("Address of v_IntSlice is: %d\n", &v_IntSlice[0])
	fmt.Printf("Address of v_AnotherIntSlice is: %d\n", &v_AnotherIntSlice[0])

	// 如果使用 `append` 追加元素的时候超出了切片对象的容量，Golang 会重新创建一个匿名数组来保存新的切片对象中的数据
	v_AnotherIntSlice = append(v_IntSlice, 100, 200)
	fmt.Println("v_AnotherIntSlice = append(v_IntSlice, 100, 200), v_IntSlice is: ", v_IntSlice)
	fmt.Println("v_AnotherIntSlice = append(v_IntSlice, 100, 200), v_AnotherIntSlice is: ", v_AnotherIntSlice)
	fmt.Printf("Address of v_IntSlice is: %d\n", &v_IntSlice[0])
	fmt.Printf("Address of v_AnotherIntSlice is: %d\n", &v_AnotherIntSlice[0])

	// 使用切片对象作为函数参数传递时是值传递，但是由于 `slice` 是引用类型，所以不会拷贝相关数组的值
	v_IntSlice = make([]int, 5, 10)
	fmt.Printf("Address of v_IntSlice is: %d\n", &v_IntSlice[0])
	sliceTest(v_IntSlice)

}
