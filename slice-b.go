package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 3, 6)
	fmt.Printf("%p\n", s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
	// test phrase 2
	b := []int{1, 2, 3, 4, 5}
	sb1 := b[2:5]
	sb2 := b[1:3]
	fmt.Println(sb1, sb2)
	sb1[0] = 9
	fmt.Println(sb1, sb2)
}
