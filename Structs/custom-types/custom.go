package main

import "fmt"

// Creating Other Custom Types & Adding Methods

type str string

func (text str) log() {
	fmt.Println(text)
}

func main(){
	var name str = "Ahmed"
	name.log()
}
