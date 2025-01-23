package main

import "fmt"

func main(){
	userNames := make([]string, 2 , 5)

	userNames[0] = "John"
	userNames[1] = "Doe"

	userNames = append(userNames, "Jane")
	userNames = append(userNames, "Smith")

	fmt.Println(userNames)

	skillRating := make(map[string]float64, 3)
	skillRating["Go"] = 4.5
	skillRating["Python"] = 4.0
	skillRating["Java"] = 3.5

	fmt.Println(skillRating)


	for index, value := range userNames {
		fmt.Println("index:", index, "value:", value)
	}

	for key, value := range skillRating {
		fmt.Println("key:", key, "value:", value)
	}
}