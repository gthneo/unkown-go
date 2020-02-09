/* a example about array
very cool
*/ 
package main

import (
	"fmt"
)

func main() {
	var a [2]int
	b := [...]int{19: 1}
	var p *[20]int = &b
	x, y := 1, 2
	c := [...]*int{&x, &y}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(p)
	p2 := new([10]int)
	fmt.Println(p2)
}
