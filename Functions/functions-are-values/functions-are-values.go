package main

import "fmt"

type transFromFunc func(int) int

func main(){
	numbers := []int{1,2,3}
	ages:= []int{10,20,30}
	doubleNumbers := transFrom(&numbers, double)
	tripleNumbers := transFrom(&numbers, triple)
	fmt.Println(doubleNumbers)
	fmt.Println(tripleNumbers)

	transFromFunc1 := getTransfromFunc("double")
	transFromFunc2 := getTransfromFunc("triple")

	transfromNumbers1 := transFrom(&numbers, transFromFunc1)
	transfromNumbers2 := transFrom(&ages, transFromFunc2)

	fmt.Println(transfromNumbers1)
	fmt.Println(transfromNumbers2)
	
}

func transFrom(numbers *[]int, transfrom transFromFunc) []int {
	transFromed := []int{}
	for _, vlaue := range *numbers {
		transFromed = append(transFromed, transfrom(vlaue))
	}
	return transFromed
}

func getTransfromFunc(name string) transFromFunc {
	if name == "double" {
		return double
	}
	return triple
}

func double(n int) int {
	return n * 2
}

func triple(n int) int {
	return n * 3
}