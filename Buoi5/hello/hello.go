package main

import (
	"fmt"
	"hello/demo"
	"hello/hi"
)

func main() {
	v := hi.GetValue()
	a := hi.GetValue2()
	demo.Demo()
	fmt.Println(hi.Name)
	fmt.Println(v)
	fmt.Println(a)

}
