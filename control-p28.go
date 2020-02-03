package main

import (
	"fmt"
)

func main() {
	a := 10
	var p *int = &a
	if a := 1; a > 0 {
		fmt.Println(a)
	}
	fmt.Println(a)
	// for statement: First style
	b := 1
	for {
		b++
		if b > 3 {
			break
		}
		fmt.Println(b)
		fmt.Println(p)
		fmt.Println(a)
	}
	fmt.Println("Over 1 style")
	// the for statement : Second style
	b = 1
	for b <= 3 {
		b++
		fmt.Println(b)
	}
	fmt.Println("Over 2 style")
	// the for statement : Third style
	a = 2
	for i := 0; i < 3; i++ {
		a++
		fmt.Println(i)
		fmt.Println(a)
	}
	fmt.Println("Over 3 style")
	//swith statement
	fmt.Printf("Swtich statement: first Style")
	a = 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}
	// swith statement: second style
	a = 1
	switch {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("a>=1")
	default:
		fmt.Println("None")
	}
	//swith statement: third style
	switch a := 1; {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("a>=1")
	default:
		fmt.Println("None")
	}
}
