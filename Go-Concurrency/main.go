package main

import (
	"fmt"
	"time"
)

func doWork(done <- chan bool){
	for {
		select {
		case <- done:	
			return
		default:
			fmt.Println("Working")
		}
	}
}

func main(){
	done := make(chan bool)

	go doWork(done)
	time.Sleep(2 * time.Second)

	close(done)

}

// func main(){
// 	bufferedChannel := make(chan string, 3)
// 	names := []string {"Ahmed", "Mohammed", "Ali"}

// 	for _, name := range names {
// 		select {
// 		case bufferedChannel <- name:
// 			fmt.Println("Sent", name)
// 		}
// 	}
// 	close(bufferedChannel)
// 	for name := range bufferedChannel {
// 		fmt.Println("Received", name)
// 	}
// }