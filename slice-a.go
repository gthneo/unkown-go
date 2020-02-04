package main

import (
	"fmt"
)

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	s1 := a[5:10]
	fmt.Println(s1)
	s2 := a[5:]
	fmt.Println(s2)
	s3 := a[:5]
	fmt.Println(s3)
	s4 := make([]int, 3, 10)
	fmt.Println(len(s4), cap(s4))
	s4 = a[0:2]
	fmt.Println(s4)
}
