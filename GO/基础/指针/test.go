package main

import (
	"time"
	"unsafe"
)

type data struct {
	x [1024 * 100]byte
}

func test() uintptr {
	p := &data{}
	return uintptr(unsafe.Pointer(p))
}

func main() {
	const N = 10000
	cache := new([N]uintptr)
	for i := 0; i < N; i++ {
		cache[i] = test()
		time.Sleep(time.Millisecond)
	}
}

//http://wiki.jikexueyuan.com/project/the-way-to-go/10.8.html
//10.8 垃圾回收和 SetFinalizer
//如果需要在一个对象 obj 被从内存移除前执行一些特殊操作，比如写到日志文件中，可以通过如下方式调用函数来实现：
//runtime.SetFinalizer(obj, func(obj *typeObj))
//在对象被 GC 进程选中并从内存中移除以前，SetFinalizer 都不会执行，即使程序正常结束或者发生错误
