package main

import (
	"math"
	"sync"
)

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}

	println(id, x)
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			sum(id)
		}(i)
	}

	wg.Wait()
}

/**
执行
go build -o test.exe
time -p ./test

0 9223372030412324865
1 9223372030412324865

real 7.70 // 程序开始到结束时间差 (⾮非 CPU 时间)
user 7.66 // ⽤用户态所使⽤用 CPU 时间⽚片 (多核累加)
sys 0.01 // 内核态所使⽤用 CPU 时间⽚片

$ GOMAXPROCS=2 time -p ./test

0 9223372030412324865
1 9223372030412324865

real 4.18
user 7.61 // 虽然总时间差不多，但由 2 个核并⾏行，real 时间⾃自然少了许多。
sys 0.02

*/
