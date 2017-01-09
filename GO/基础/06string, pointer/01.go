package main

import "fmt"

func main() {

	// 默认 string 类型对象零值为空字符串，尾部不包含 `\0`
	var emptyStr string
	fmt.Println("emptyStr is: ", emptyStr)
	fmt.Println("len(emptyStr) is: ", len(emptyStr))

	// 声明 string 对象并初始化，使用下标访问。
	// 注意，如果字符串中包含中文等非 ASIC 码的字符，则使用下标索引会导致与期望不符的结果
	// 字符串是 UTF-8 编码，所以非 ASIC 码字符不止一个字节，而使用下标获得的是每个字节的内容。
	str := "I like 高圆圆"
	fmt.Println("string object is: ", str)
	fmt.Println("len(str) = : ", len(str))
	fmt.Println("str[1]: ", str[1])

	// 使用 ` 语法可以声明无转义的字符串
	str = `I
		Love
		Golang`

	fmt.Println("str: ", str)

	// 修改字符串，注意将字符串分别转化为 `rune` 和 `byte` 的 `slice` 对象。
	// 它们的长度是不相等的，`rune` 切片对象的长度等于原始 `string` 对象中的字符个数。
	// `byte` 切片对象的长度等于原始 `string` 对象在内存中所占的字节数，其值和 `len` 取得的 `string` 对象长度相等
	// 如果 `string` 对象中有非 ASIC 码，则两者的长度是不相等的。
	// 所以为了保证兼容性，最好将 `string` 对象转化为 `rune` 切片对象后进行修改。
	str = "I like 高圆圆"
	fmt.Println("Before modification, str: ", str)
	fmt.Println("len(str) = : ", len(str))
	rune_str := []rune(str)
	byte_str := []byte(str)
	fmt.Println("len([]rune(str)) = ", len(rune_str))
	fmt.Println("len([]byte(str)) = ", len(byte_str))
	rune_str[7] = '范'
	rune_str[8] = '冰'
	rune_str[9] = '冰'
	fmt.Println("After modification, str:  ", string(rune_str))

	// 单引号中的字符常量为 `rune` 类型，取类型后为 int32
	v_char := 'g'
	fmt.Printf("The type of v_char is %T\n", v_char)

}
