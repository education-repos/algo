package main

import "fmt"

/*
Дан массив A[1...n] из n различных натуральных чисел, в порядке возрастания
Дан второй массив B[1...n] из n натуральных чисел
Для каждого элемента массива B вывести позицию в массиве A, если он есть
или -1 если его нет. Результат представить в виде массива

A[1...n] = 1 5 8 12 13
B[1...n] = 8 1 23 1 11
result = 3 1 -1 1 -1
*/

// findIndexes ищет индексы вхождений массива B в массив А
func findIndexes(inputA, inputB []int) []int {
	output := make([]int, 0, len(inputA))

	for _, v := range inputB {
		if position, ok := findPositon(inputA, v); ok {
			output = append(output, position+1)
		} else {
			output = append(output, -1)
		}

	}

	return output
}

// findPositon ищет позицию элемента v в массиве A используя двоичный поиск
func findPositon(inputA []int, v int) (int, bool) {
	left := 0
	right := len(inputA) - 1
	for left <= right {
		pos := (left + right) / 2
		if inputA[pos] == v {
			return pos, true
		} else if inputA[pos] > v {
			right = pos - 1
		} else {
			left = pos + 1
		}
	}

	return -1, false
}

func main() {
	inputA := []int{1, 5, 8, 12, 13}
	inputB := []int{8, 1, 23, 1, 11}

	result := findIndexes(inputA, inputB)

	fmt.Printf("array A: %v\narray B: %v \nresult: %v\n", inputA, inputB, result)
}
