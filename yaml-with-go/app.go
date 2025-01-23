package main

import (
	"fmt"
	"os"
)

func ReadFile() string{
	data, _ := os.ReadFile("config.yaml")
	dataText := string(data)
	return dataText
}

func main(){
	data := ReadFile()
	fmt.Println(data)
}