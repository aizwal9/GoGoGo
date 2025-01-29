package main

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}	

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64{
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * 3.14 * c.radius
}

func measure(g geometry) (float64,float64){
	return g.area(),g.perim()
}

func detectCircle(g geometry) {
	if c,ok := g.(circle); ok {
		fmt.Println("Circle detected", c)
	}
}

func main(){
	r := rect{width: 3,height: 4}
	c := circle{radius: 5}

	a1,p1 := measure(r)
	a2,p2 := measure(c)

	fmt.Println("Area of rectangle: ",a1)
	fmt.Println("Perimeter of rectangle: ",p1)

	fmt.Println("Area of circle: ",a2)
	fmt.Println("Perimeter of circle: ",p2)
	detectCircle(r)
	detectCircle(c)
}