package main

import(
	"fmt"
)

func main(){
	var m map[int]string
	m = make(map[int]string)
	m[1]="OK"
	m[2]="OK2"
	a := m[1]
	fmt.Println(a)
	fmt.Println(m)
	
}