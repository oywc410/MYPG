package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	ID      int
	Name    string
	Email   string
	Age     int
	Address string
	memo    string
}

//定义json格式
type Person2 struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"-"`
	Age     int    `json:"age"`
	Address string `json:"address,omitempty"`
	memo    string
}

func main() {

	//编码
	person := &Person{
		ID:      1,
		Name:    "aaa",
		Email:   "bbbb",
		Age:     5,
		Address: "",
		memo:    "",
	}
	b, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	person2 := &Person2{
		ID:      1,
		Name:    "aaa",
		Email:   "bbbb",
		Age:     5,
		Address: "",
		memo:    "",
	}
	b2, err2 := json.Marshal(person2)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(string(b2))

	//解码
	var person3 Person
	b3 := []byte(`{"id":1, "name":"Gopher", "age":5}`)
	err3 := json.Unmarshal(b3, &person3)
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println(person3)
}
