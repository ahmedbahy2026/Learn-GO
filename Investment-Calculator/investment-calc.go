package main

import (
	"fmt"
	"math"
)

func main(){
	const inflationRate = 2.5
	var investmentAmout, years float64 = 1000, 10
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmout)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmout * math.Pow((1 + expectedReturnRate/100), float64(years));

	futureRealValue := futureValue/ math.Pow(1 + inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}