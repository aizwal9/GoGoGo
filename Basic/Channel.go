package main

import "fmt"

func main(){

	messages := make(chan string)

	go func(){
		messages <- "ping"
	}()

	// msg := <- messages
	fmt.Println(<- messages)

	// Buffered

	channel := make(chan string,2)

	channel<- "buffered"
	channel<- "channel"
	fmt.Println(<- channel)
	fmt.Println(<- channel)


}