package main

import "fmt"


func main(){
	numbers := []int{1,2,3,4,5}
	fmt.Println(sum(numbers...))
	fmt.Println(sum(1,2,3,4,5))
	fmt.Println(sum([]int{1,2,3,4,5}...))


}

func sum(nums ...int)int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}