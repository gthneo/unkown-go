package main
import (
	"fmt"
)
type person struct{
	Name string
	Age int
}

func main() {
	a := person{}
	a.Name = "Joe"	
	a.Age = 19
	fmt.Println(a)
}