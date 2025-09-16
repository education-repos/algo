package main

import (
	"fmt"
	"time"
)

/*
Если необходимо ввести проверку на переполнение Int можно использовать math.MaxInt32/math.MaxInt64
необходимо проверять до суммирования, что a не больше чем max-b.
Второй вариант - использовать библиотеку safemath.CheckedAddU64 (Go 1.21+)
*/

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

// CalculateFibonacciWithCache кеширует результаты промежуточных вычислений в мапе
func CalculateFibonacciWithCache(n int, cache map[int]int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return n
	}

	if v, ok := cache[n]; ok {
		return v
	}

	result := CalculateFibonacciWithCache(n-1, cache) + CalculateFibonacciWithCache(n-2, cache)
	cache[n] = result

	return result
}

// CalculateFibonaciWithoutRecursion решение без рекурсии, но с использованием кеширующего слоя
func CalculateFibonaciWithoutRecursion(n int) int {
	if n < 0 { // проверяем только ноль
		return 0
	}

	cache := make(map[int]int, n) // мапка под кеш
	cache[0], cache[1] = 0, 1     // заполняем первые два элемента, чтобы считать остальные

	for i := 2; i <= n; i++ { // со второго и до n включительно
		cache[i] = cache[i-1] + cache[i-2] // текущий элемент равен сумме двух предыдущих
	}

	return cache[n] // последний элемент для n
}

// CalculateFibonaciOnArray решение без рекурсии, но с использованием кеширующего слоя на массиве
// можем использовать слайс(массив) т.к. возрастающая последовательность целых чисел
func CalculateFibonaciOnArray(n int) int {
	if n < 0 { // проверяем только ноль
		return 0
	}

	cache := make([]int, n+1) // мапка под кеш
	cache[0], cache[1] = 0, 1 // заполняем первые два элемента, чтобы считать остальные

	for i := 2; i <= n; i++ { // со второго и до n включительно
		cache[i] = cache[i-1] + cache[i-2] // текущий элемент равен сумме двух предыдущих
	}

	return cache[n] // последний элемент для n
}

// CalculateFibonaciOnVars решение без рекурсии, используя лишь 2 переменные
// т.к. нам достаточно хранить только 2 предыдущих значения, то можно избавиться от кеша
func CalculateFibonaciOnVars(n int) int {
	if n < 0 { // проверяем только ноль
		return 0
	}

	if n <= 1 {
		return n
	}

	prev, current := 0, 1 // заполняем предыдущий и текущий элементы

	for i := 2; i <= n; i++ { // со второго и до n включительно
		// предыдущий равен текущему, а новый текущий равен сумме предыдущего и старого текущего
		prev, current = current, prev+current
	}

	return current // последний текущий
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
	input := 45

	startTime := time.Now()
	output := CalculateFibonacci(input)
	timing1 := time.Since(startTime)
	fmt.Printf("=====\nFibonacci for %d: %v\nTime: %0.fm %0.fs %dms\n", input, output,
		timing1.Minutes(), timing1.Seconds(), timing1.Microseconds())

	cache := make(map[int]int, input)
	startTimeWithCache := time.Now()
	output = CalculateFibonacciWithCache(input, cache)
	timing2 := time.Since(startTimeWithCache)
	fmt.Printf("=====\nFibonacci for %d: %v\nTime: %0.fm %0.fs %dms\n", input, output,
		timing2.Minutes(), timing2.Seconds(), timing2.Microseconds())

	startTimeWithoutRecursion := time.Now()
	output = CalculateFibonaciWithoutRecursion(input)
	timing3 := time.Since(startTimeWithoutRecursion)
	fmt.Printf("=====\nFibonacci for %d: %v\nTime: %0.fm %0.fs %dms\n", input, output,
		timing3.Minutes(), timing3.Seconds(), timing3.Microseconds())

	startTimeOnArray := time.Now()
	output = CalculateFibonaciOnArray(input)
	timing4 := time.Since(startTimeOnArray)
	fmt.Printf("=====\nFibonacci for %d: %v\nTime: %0.fm %0.fs %dms\n", input, output,
		timing4.Minutes(), timing4.Seconds(), timing4.Microseconds())

	startTimeOnVars := time.Now()
	output = CalculateFibonaciOnVars(input)
	timing5 := time.Since(startTimeOnVars)
	fmt.Printf("=====\nFibonacci for %d: %v\nTime: %0.fm %0.fs %dms\n", input, output,
		timing5.Minutes(), timing5.Seconds(), timing5.Microseconds())

	fmt.Printf("=====\nLast number: %d\n", GetLastNumber(output))

	fmt.Printf("=====\nModule for 2 is: %d\n", Module(output, 2))

	return
}
