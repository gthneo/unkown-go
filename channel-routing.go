package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go Go Go !!!")
		c <- true
	}()
	<-c
	d := make(chan bool)
	go func() {
		fmt.Println("Go Go Go GO!!!!")
		d <- true
		close(d)
	}()
	for v := range d {
		fmt.Println(v)
	}
}
