package main

import "fmt"

type transFromFunc func(int) int

func main(){
	numbers := []int{1,2,3}
	doubleNumbers := transFrom(&numbers, func(n int)int { return n * 2})
	tripleNumbers := transFrom(&numbers, func(n int)int { return n * 3})
	fmt.Println(doubleNumbers)
	fmt.Println(tripleNumbers)

	double := createTranformer(2)
	triple := createTranformer(3)

	doubled := transFrom(&numbers, double)
	tripled := transFrom(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)


}

func transFrom(numbers *[]int, transfrom transFromFunc) []int {
	transFromed := []int{}
	for _, vlaue := range *numbers {
		transFromed = append(transFromed, transfrom(vlaue))
	}
	return transFromed
}


func createTranformer(factor int) func(int)int {
	return func(number int) int {
		return number * factor
	} 
}