package main

import (
    "fmt"
    "runtime"
)

func main() {
    go fmt.Println("Go!")
    runtime.Gosched()
}