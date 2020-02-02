package main

import (
	"fmt"
)

const a int = 1
const b = 'A'
const a, b, c = 1,

func main() {
	var a float32
	a = 100.1
	fmt.Println(a)
	b := bool(a)
	fmt.Println(b)
	fmt.Println(a)

}
