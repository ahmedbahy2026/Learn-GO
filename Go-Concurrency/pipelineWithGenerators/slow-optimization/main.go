package main

import (
	"fmt"
	"math/rand"
	"time"
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

func take[T any, K any](done <- chan K, stream <- chan T, n int) <- chan T {
	taken := make(chan T)

	go func(){
		defer close(taken)
		for i := 0; i < n ; i++ {
			select {
			case <-done:
				return;
			case taken <- <-stream:
			}
		}
	}()

	return taken
}

func primeFinder(done <- chan int, randomStream <- chan int) <- chan int {
	isPrime := func (n int) bool {
		if n == 1 {
			return false
		}
		for i := 2 ; i < n ; i++ {
			if n % i == 0 {
				return false
			}
		} 
		return true
	}

	primes := make(chan int)

	go func() {
		defer close(primes)
		for {
			select {
			case <- done: 
				return;
			case randomInt := <-randomStream:
				if isPrime(randomInt) {
					primes <- randomInt
				}
			}
		}
	}()
	return primes
}

func main(){
	start := time.Now()
	done := make(chan int)
	randomFunc := func() int { return rand.	Intn(1e9)}
	randomStream := repeatFunc(done, randomFunc)
	primeStream := primeFinder(done, randomStream)
	found := take(done, primeStream, 10)

	defer close(done)

	for i := range found {
		fmt.Println(i)
	}

	fmt.Println(time.Since(start))
}