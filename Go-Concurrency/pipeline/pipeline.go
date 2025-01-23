package main

import "fmt"

func sliceToChannel(input []int) <- chan int {
	
	out := make(chan int)
	go func () {
		for _, i := range input {
			out <- i 
		}
		close(out)
	}()
	return out
}

func square(input <- chan int) <- chan int {
	out := make(chan int)
	go func () {
		for i := range input {
			out <- i * i
		}
		close(out)
	}()
	return out
}

func printer(input <- chan int) {
	for i := range input {
		fmt.Println(i)
	}
}

func main(){
	// Create a pipeline of 3 stages
	input := []int{1,3,5,7,9}

	// First stage is a generator function that generates 10 integers
	dataChannel := sliceToChannel(input)

	// Second stage is a squarer function that squares the integers
	squaredChannel := square(dataChannel)

	// Third stage is a printer function that prints the squared integers
	printer(squaredChannel)
}
