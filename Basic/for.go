package main

import (
	"fmt"
)

func main(){

	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	var i=0;
	for i<10{
		if i%2 ==0 {
			fmt.Println("even")
		} else {
			fmt.Println("Odd")
		}
		i++;
	}

	switch i {
	case 0: fmt.Println("Zero")
	case 10 : fmt.Println("Ten")
	}	
}