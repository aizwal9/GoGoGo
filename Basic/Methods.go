package main

import "fmt"

type rectangle struct {
	lenght,breadth int
}

func (r rectangle) area() int{
	return r.lenght * r.breadth
}

func main(){
	r := rectangle{lenght: 10,breadth: 5}
	fmt.Println("Area: ",r.area())
	fmt.Println("Rectangle: ",r.area())

}