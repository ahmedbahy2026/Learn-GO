package main

import "fmt"

func main() {
	// Creating a channel (Channels are FIFO Queue) 
	myChannel := make(chan string)

	go func() {
		myChannel <- "Hello"
		fmt.Println("Sent Hello")
		myChannel <- "World"
		fmt.Println("Sent World")
		
	}()

	fmt.Println("Waiting for messages")
	msg := <-myChannel 
	fmt.Println("Received", msg)
	msg += " "
	fmt.Println("Waiting for messages")
	msg += <- myChannel
	fmt.Println("Received", msg)

	fmt.Println(msg)
}