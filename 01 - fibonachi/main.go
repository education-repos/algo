package main

import (
	"fmt"
	"time"
)

// CalculateFibonacci необходимо найти число фибоначчи
func CalculateFibonacci(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return n
	}

	return CalculateFibonacci(n-1) + CalculateFibonacci(n-2)
}

// GetLastNumber необходимо найти последнюю цифру числа фибоначчи
func GetLastNumber(n int) int {
	return n % 10
}

// Module необходимо найти остаток от деления n-го числа фибоначчи на m
func Module(n, m int) int {
	return n % m
}

func main() {
	startTime := time.Now()
	input := 10
	output := CalculateFibonacci(input)
	timing := time.Since(startTime)

	fmt.Printf("Fibonacci for %d: %v\nTime: %0.fm %0.fs %dms\n", input, output,
		timing.Minutes(), timing.Seconds(), timing.Microseconds())

	fmt.Printf("Last number: %d\n", GetLastNumber(output))

	fmt.Printf("Module for 2 is: %d\n", Module(output, 2))

	return
}
