package main

import (
	"fmt"
	"time"
)

func print(str string){
	fmt.Println(str)
}

func main(){
	go print("Hello")
	go print("World")

	fmt.Println("Done")
	time.Sleep(2 * time.Second)
}