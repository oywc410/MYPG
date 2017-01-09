package main

import "fmt"

type Person struct {
    Name    string
	Gender  string
	Age     uint8
    Address string
}

//* 可以直接访问更改
func(person *Person) Move(newAddress string)string{
    old:=person.Address
    person.Address=newAddress
    return old
}

func main() {
	//省略写法  也可以写为 {Name : "Robert", Gender : "Male" ...}
	p := Person{"Robert", "Male", 33, "Beijing"}
	oldAddress := p.Move("San Francisco")
	fmt.Printf("%s moved from %s to %s.\n", p.Name, oldAddress, p.Address)
}