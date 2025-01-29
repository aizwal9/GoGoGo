package main

import (
	"fmt"
	"time"
)

func f(s string){
	for i:=0;i<3;i++ {
		fmt.Println(s,":", i)
	}
}

func main(){

	f("direct")

	go f("goroutine")
	
	go func (msg string)  {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	println("done")

}