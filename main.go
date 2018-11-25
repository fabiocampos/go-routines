package main

import (
	"fmt"
	"time"
)

type Fibonacci struct {
	FibonacciRequest  int
	FibonacciResponse int
}

func main() {
	//CallFibonacciSync()
	CallFibonacciAsync()
}

func CallFibonacciAsync() {
	start := time.Now()
	numbers := []int{30, 50, 40, 52}
	channel := make(chan Fibonacci, len(numbers))

	for _, fibonacciToDiscover := range numbers {
		go AsyncFibonacciService(fibonacciToDiscover, channel)
	}

	for i := 0; i < len(numbers); i++ {
		fib := <-channel
		fmt.Printf("O fibonacci de %d é %d \n", fib.FibonacciRequest, fib.FibonacciResponse)
	}
	close(channel)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func CallFibonacciSync() {
	start := time.Now()
	numbers := []int{30, 50, 40, 52}
	for _, fibonacciToDiscover := range numbers {
		fib := CalcFibonacci(fibonacciToDiscover)
		fmt.Printf("O fibonacci de %d é %d \n", fibonacciToDiscover, fib)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func AsyncFibonacciService(fibonacciToDiscover int, response chan<- Fibonacci) {
	fibonacci := CalcFibonacci(fibonacciToDiscover)
	response <- Fibonacci{FibonacciRequest: fibonacciToDiscover, FibonacciResponse: fibonacci}
}

func CalcFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return CalcFibonacci(n-1) + CalcFibonacci(n-2)
}
