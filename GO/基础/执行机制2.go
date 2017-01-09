package main

import (
	"fmt"
	"reflect" //返射 http://golanghome.com/post/105
	"unsafe"
)

type iface struct {
	itab, data uintptr
}

func main() {

	//http://my.oschina.net/goal/blog/194233

	//interface 由 type data组成
	//var a interface{} = int(1)时 其实为 {int, 1}
	//只有 tab 和 data 都为 nil 时，接⼝口才等于 nil
	var a interface{} = nil
	var b interface{} = (*int)(nil)

	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))

	fmt.Println(a == nil, ia)
	fmt.Println(b == nil, ib, reflect.ValueOf(b).IsNil())
	/**
	true {0 0}
	false {5013504 0} true
	*/

	var val interface{} = (*interface{})(nil)
	// val = (*int)(nil)
	if val == nil {
		fmt.Println("val is nil")
	} else {
		fmt.Println("val is not nil")
	}

	//val is not nil
}
