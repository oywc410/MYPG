package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	ID      int
	Name    string
	Email   string
	Age     int
	Address string
	memo    string
}

func main() {
	setJsonFile()
	getJsonFile()
}

func setJsonFile() {
	person := &Person{
		ID:      1,
		Name:    "aaa",
		Email:   "bbbb",
		Age:     5,
		Address: "",
		memo:    "",
	}

	file, err := os.Create("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		fmt.Println(1)
		file.Close()
	}()

	//获取解码数据
	encoder := json.NewEncoder(file)
	err = encoder.Encode(person)

	if err != nil {
		log.Fatal(err)
	}
}

func getJsonFile() {
	file, err := os.Open("./person.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(2)

	defer func() {
		fmt.Println(3)
		file.Close()
	}()

	var person Person

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(person)
}
