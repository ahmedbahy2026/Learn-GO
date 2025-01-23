package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func process(data int) int {
	time.Sleep(2 * time.Second)
	return 2 * data
}

func processDate(wg * sync.WaitGroup, result * []int, data int) {
	defer wg.Done()
	processedData := process(data)
	lock.Lock()
	*result = append(*result, processedData)
	lock.Unlock()
}

func main(){
	start := time.Now() 

	input := []int{1,2,3,4,5}
	result := []int{}

	var wg sync.WaitGroup
	for _, val := range input {
		wg.Add(1)
		go processDate(&wg, &result, val)
	}
	wg.Wait()

	fmt.Println("Time Elapsed: ", time.Since(start))
	fmt.Println(result)
}