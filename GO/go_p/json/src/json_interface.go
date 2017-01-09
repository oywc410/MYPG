package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var f interface{}
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	json.Unmarshal(b, &f)

	fmt.Println(f) //不可访问实际输出

	//fmt.Println(f["Name"]) 报错

	m := f.(map[string]interface{}) //通过断言来访问数据
	fmt.Println(m)

	for k, v := range m {
		switch w := v.(type) {
		case string:
			fmt.Println(k, "is string", w)
		case int:
			fmt.Println(k, "is int", w)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range w {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	fmt.Println(m["Name"])
}
