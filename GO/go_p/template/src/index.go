package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	Email    string
}

func main() {
	t := template.New("filename example")
	t, _ = t.Parse("hello {{.UserName}} {{.Email}}!")
	p := Person{UserName: "Astaxie"}
	t.Execute(os.Stdout, p)
}
