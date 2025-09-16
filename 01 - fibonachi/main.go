package main

import (
	"fmt"
	"time"
)

func CalculateFibonacci(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return n
	}

	return CalculateFibonacci(n-1) + CalculateFibonacci(n-2)
}

func main() {
	startTime := time.Now()
	input := 40
	output := CalculateFibonacci(input)
	timing := time.Since(startTime)

	fmt.Printf("Fibonacci for %d: %v\nTime: %0.fm %0.fs %dms\n", input, output,
		timing.Minutes(), timing.Seconds(), timing.Microseconds())

	return
}
