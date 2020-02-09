package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	a := &person{
		Name: "Joe",
		Age:  19,
	}
	a.Name = "Ok"
	fmt.Println(a)
	A(a)
	B(a)
	fmt.Println(a)
}

func A(per *person) {
	per.Age = 13
	fmt.Println("A", per)
}

func B(per *person) {
	per.Age = 15
	fmt.Println("A", per)
}
