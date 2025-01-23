package main

import (
	"fmt"
	"sync"
	"time"
)

func process(data int) int {
	time.Sleep(2 * time.Second)
	return 2 * data
}

func processDate(wg * sync.WaitGroup, result *int, data int) {
	defer wg.Done()
	processedData := process(data)
	*result = processedData
}

func main(){
	start := time.Now() 

	input := []int{1,2,3,4,5}
	result := make([]int, len(input))

	var wg sync.WaitGroup
	for i, val := range input {
		wg.Add(1)
		go processDate(&wg, &result[i], val)
	}
	wg.Wait()

	fmt.Println("Time Elapsed: ", time.Since(start))
	fmt.Println(result)
}