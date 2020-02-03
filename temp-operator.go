package main

import (
	"fmt"
)

const (
	B          = 1
	KB float64 = 1 << (iota * 10)
	MB
	GB
	PB
)

func main() {
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(PB)
}
