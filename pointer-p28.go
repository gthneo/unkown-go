// not support 指针运算
// & 取变量地址
// * 通过指针间接访问目标对象。
// 默认值是nil， 而非NULL
package main

import (
	"fmt"
)

func main() {
	a := 1
	a++
	fmt.Println(a)
	a++
	fmt.Println(a)
	var p *int
	fmt.Println(p)
	p = &a
	fmt.Println(*p)
	fmt.Println(p)
}
