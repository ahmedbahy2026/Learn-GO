package main

import (
	"fmt"
	"math/rand"
)

func repeatFunc[T any, K any](done <- chan K, fn func()T) <- chan T{
	stream := make(chan T)
	
	go func(){
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()
	return stream
}

func main(){
	done := make(chan int)
	randomFunc := func() int { return rand.Intn(1e9)}

	defer close(done)

	for i := range repeatFunc(done, randomFunc) {
		fmt.Println(i)
	}


}