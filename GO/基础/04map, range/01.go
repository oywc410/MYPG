package main

import "fmt"

func main() {
	// 如果已经知道 map 中的数据，可以直接以下面的方式声明 map 对象
	v_map := map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
	}
	fmt.Println("v_map: ", v_map)
	fmt.Println("len(v_map): ", len(v_map))

	// 从下面的用法可以看到 “修改已有值” 和 “添加新 pair” 形式上是一样的。
	// 如果 key 值已存在，就是修改操作，否则就是添加新 pair
	// 修改已有值
	v_map[1] = "One hundred"

	// 添加新 pair
	v_map[4] = "Two hundred"
	fmt.Println("After changes, v_map: ", v_map)
	fmt.Println("After changes, len(v_map): ", len(v_map))

	// 如果只是声明一个 map 对象，可以使用内置函数 make
	v_mapByMake := make(map[int]string)
	v_mapByMake[0] = "One hundred"
	fmt.Println("make(map[int]string, 10)，v_mapByMake: ", v_mapByMake)

	// 使用 make 时也可以设定预期的键值对数量，在初始化时一次性分配大量内存，从而避免使用过程中频繁动态分配
	// 这里给定的数量值不会影响初始化后 len(mapObject)
	v_mapByMake = make(map[int]string, 10)
	fmt.Println("make(map[int]string, 10)，Before chagnes，v_mapByMake: ", v_mapByMake)
	fmt.Println("make(map[int]string, 10)，Before chagnes，len(v_mapByMake): ", len(v_mapByMake))
	v_mapByMake[0] = "One hundred"
	v_mapByMake[1] = "Two hundred"
	v_mapByMake[2] = "Three hundred"
	fmt.Println("make(map[int]string, 10)，After chagnes，v_mapByMake: ", v_mapByMake)

	// 遍历 map 对象中的键值对
	// 注意，这里迭代的顺序是不确定的
	for key, value := range v_mapByMake {
		fmt.Printf("%d : %s\n", key, value)
	}

	// 检查 map 对象中是否存在某个 key 索引的元素，如果存在获取该 key 索引的 value
	if value, ok := v_mapByMake[1]; ok {
		fmt.Println("value, ok := v_mapByMake[1]，value: ", value)
	}

	// 这里如果只是判断是否存在，可以使用占位符
	_, ok := v_mapByMake[1]
	fmt.Println("_, ok := v_mapByMake[1]，ok: ", ok)

	// 如果尝试获取不存在的元素，会返回空，不会抛出异常
	fmt.Println("v_mapByMake[10]: ", v_mapByMake[10])

	// 如果尝试删除不存在的元素，对已有数据不会有影响，不会抛出异常
	fmt.Println("Before deletion non-existed elem, v_mapByMake: ", v_mapByMake)
	delete(v_mapByMake, 10)
	fmt.Println("After deletion non-existed elem, v_mapByMake: ", v_mapByMake)

	// 从 map 中获取的 value 是原始数据的拷贝，如果其本身是值类型，对其修改时不允许的。
	// 如果是引用类型，修改就没有问题。
	// 比如下面的例子，如果值是 slice 对象，对其元素的修改是允许的。
	// 如果值是 array 对象，则通过 map 的索引获取到的 value 是不允许修改的。
	// 将下面 map 的元素类型修改为 array 类型 map[int][3]int, 尝试修改会导致编译错误
	v_mapOfArray := map[int][]int{
		1: {0, 1, 2},
		2: {3, 4, 5},
	}
	fmt.Println("Before change, v_mapOfArray is: ", v_mapOfArray)
	fmt.Println("Before change, v_mapOfArray[1][0] is: ", v_mapOfArray[1][0])
	v_mapOfArray[1][0] = 100
	fmt.Println("After change, v_mapOfArray[1][0] is: ", v_mapOfArray[1][0])
	fmt.Println("After change, v_mapOfArray is: ", v_mapOfArray)

	v_map = map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
	}

	fmt.Println("Before iterating, v_map: ", v_map)
	// 迭代过程中可以安全地删除键值对，在迭代过程中也支持添加新的键值对
	for key, _ := range v_map {
		delete(v_map, key)
		if key == 1 {
			v_map[key*5] = "New"
		}
	}
	fmt.Println("After iterating, v_map: ", v_map)

}
