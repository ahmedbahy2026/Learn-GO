package main

import "fmt"

func main() {
	var prices []int
	prices = []int{1,3,10,5}
	fmt.Println(prices)
	fmt.Println(len(prices), cap(prices)) // 4 4	

	prices = append(prices, 15)
	fmt.Println(prices)
	fmt.Println(len(prices), cap(prices)) // 5 8

	prices = append(prices, 20)
	fmt.Println(prices)
	fmt.Println(len(prices), cap(prices)) // 6 8

	prices = append(prices, 25)
	fmt.Println(prices)
	fmt.Println(len(prices), cap(prices)) // 7 8

	prices = append(prices, 30)
	fmt.Println(prices)
	fmt.Println(len(prices), cap(prices)) // 8 8

	prices = append(prices, 35)
	fmt.Println(prices)
	fmt.Println(len(prices), cap(prices)) // 9 16
}




// func main() {
// 	prices := [4]int{1,3,10,5}
// 	fmt.Println(prices)
// 	fmt.Println(prices[0])

// 	var productNames [4]string
// 	productNames = [4]string{"Book", "Shoes", "Laptop", "Phone"}
// 	fmt.Println(productNames)

// 	// Selecting Parts of Arrays With Slices
// 	// slices are a reference (pointers) to a part of an array (not a copy)

// 	// [start:stop]  stop is not included (exclusive)
// 	// [:stop]  start is 0
// 	// [start:]  stop is the end of the array
// 	// [:]  start is 0 and stop is the end of the array

// 	featuredPrices := prices[1:3] // [3,10]  [start:stop]  stop is not included (exclusive)
// 	fmt.Println(featuredPrices)

// 	highlightedPrices := featuredPrices[:1] // [1,3]  [:stop]  start is 0
// 	fmt.Println(highlightedPrices)

// 	featuredPrices[0] = 100
// 	fmt.Println(prices) // [1,100,10,5]

// 	fmt.Println(len(prices), cap(prices)) // 4 4

// }