package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	numbers := []int{50, 60, 70, 80}
	for _, fibonacciToDiscover := range numbers {
		fibonacci := CalcFibonacci(fibonacciToDiscover)
		fmt.Printf("O fibonacci de %d é %d \n", fibonacciToDiscover, fibonacci)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func CalcFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return CalcFibonacci(n-1) + CalcFibonacci(n-2)
}
