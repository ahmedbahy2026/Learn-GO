package main

import "fmt"

type Product struct {
	id int
	price float64
	title string
}

func main(){
	hoppies := [3]string{"Reading", "Coding", "Traveling"}
	fmt.Println(hoppies)
	fmt.Println(hoppies[0])
	fmt.Println(hoppies[1:])

	featuredHoppies := hoppies[:2]
	// featuredHoppies := hoppies[0:2]
	fmt.Println(featuredHoppies)

	featuredHoppies = featuredHoppies[1:3]
	fmt.Println(featuredHoppies)

	goals := []string{"Goal1", "Goal2"}
	fmt.Println(goals)
	fmt.Println(goals[1]);
	goals = append(goals, "Goal3")
	fmt.Println(goals)

	products := []Product{{id: 1, price: 10.5, title: "Book"}, {id: 2, price: 20.5, title: "Shoes"}}
	// products := []Product{Product{id: 1, price: 10.5, title: "Book"}, Product{id: 2, price: 20.5, title: "Shoes"}}
	
	fmt.Println(products)
	
	products = append(products , Product{id: 3, price: 30.5, title: "Laptop"})
	fmt.Println(products)
}