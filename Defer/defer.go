package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func readFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err;
	}
	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return err;
	}

	fmt.Println(string(data))
	return nil
}

func processData(data []int) {
	start := time.Now()

	defer func() {
		fmt.Println("Time taken: ", time.Since(start))
	} ()

	for _, i := range data {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func safeOperation() {
	defer func() {
		if r := recover(); r != nil  {
			fmt.Println("Recovered from ", r)
		}
	}()

	panic("Something went wrong")
	fmt.Println("This will not be executed")
}

// Defer statements are executed in LIFO order
func moreDefer() {
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
}

func main(){
	err := readFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	data :=[]int{1,3,5,7,9}
	processData(data)

	safeOperation();

	moreDefer()

	// Defer statements are executed in LIFO order
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	for _, day := range days {
		defer fmt.Println(day)
	}
}