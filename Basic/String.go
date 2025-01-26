package main

import (
	"fmt"
)

func main(){

	var a []string
    fmt.Println("uninit:", a, a == nil, len(a) == 0)

	a = make([]string, 5)
	fmt.Println("init:", a, a == nil, cap(a))
	a[0] = "Hello"
	a[1] = "World"
	a[2] = "!"
	a = append(a,"ABC")
	fmt.Println("append:", a, cap(a), len(a))

	c := make([]string, len(a))
	copy(c, a)
	fmt.Println("copy:", c)
	fmt.Println("slice:", a[1:2])
}