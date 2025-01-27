package main

import "fmt"

func incrementVal(i int){
	i++;
}

func incrementPtr(i *int){
	*i++;
}


func main(){
	i := 7
	fmt.Println(i)
	incrementVal(i)
	fmt.Println(i)
	incrementPtr(&i)
	fmt.Println(i)
	fmt.Println(&i)
}