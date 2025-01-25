package main

import (
	"fmt"
)

func main(){
	var b [10]int
	fmt.Println(b)

	b[0] = 1
	b[1] = 2
	fmt.Println(b)

	fmt.Println(b[0])

	b = [...]int{1,2,3,4,7:5,9:6}
	fmt.Println(b)

	var twod [2][3]int
	fmt.Println(twod)
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twod[i][j] = i+j
		}
	}
	fmt.Println(twod)
}