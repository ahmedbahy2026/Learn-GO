package main

import "fmt"

func main(){
	age := 50
	
	fmt.Println(age)
	decreaseBy10(&age)
	fmt.Println(age)
}

func decreaseBy10(age *int){
	*age -= 10;
}
