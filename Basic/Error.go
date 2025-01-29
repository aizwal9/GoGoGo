package main

import (
	"errors"
	"fmt"
)

func f(i int) (int,error) {
	if(i== 10){
		return -1,errors.New("Cant work with 10")
	}
	return i+13,nil
}

var ErrOutOfTea = fmt.Errorf("Out of Tea")
var ErrPower = fmt.Errorf("No Power")

func makeTea(i int) error{
	if(i==2){
		return ErrOutOfTea
	} else if(i==4){
		return fmt.Errorf("Oh no: %w",ErrPower)
	}
	return nil
}

func main(){
	
	for _,i := range []int{7,47}{
		
		if r,e := f(i); e!=nil{
			fmt.Println("f failed :",e)
		} else{
			fmt.Println("f Passed :",r)
		}
	}

	for i := range 5{

		if err := makeTea(i); err!=nil{
			if errors.Is(err,ErrOutOfTea){
				fmt.Println("Get some tea")
			} else if errors.Is(err,ErrPower) {
				fmt.Println("Turn on Inverter")
			} else{
				fmt.Printf("unkown error %s\n",err)
			}
			continue
		}

		fmt.Println("Tea is ready")

	}
}