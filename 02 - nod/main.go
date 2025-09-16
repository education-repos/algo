package main

import (
	"fmt"
	"time"
)

// BadAlgoForGCD медленный алгоритм для поиска наименьшего общего делителя с перебором всех значений
func BadAlgoForGCD(a, b int) int {
	if a <= 0 || b <= 0 {
		return 0
	}

	current := 1
	// идем последовательно от 1 до большего из a и b
	for d := 1; d <= a && d <= b; d++ {
		if a%d == 0 && b%d == 0 { // проверяем что оба числа делятся на делитель
			if d > current { // проверяем что текущий делитель больше предыдущего сохраненного
				current = d // сохраняем
			}
		}
	}

	return current
}

// MoreEffectiveGCD ищем делитель через a=d1*d2, которые min(d1, d2) <= sqrt(a) <= max(d1, d2), d2 = a/d1
func MoreEffectiveGCD(a, b int) int {
	if a <= 0 || b <= 0 {
		return 0
	}

	current := 1
	// перебираем все маленькие делители
	for d := 1; d*d <= a; d++ { // выполняется условие, что делитель не превосходит корень а (квадрат а)
		if a%d == 0 { // является делителем a
			if b%d == 0 { // является делителем b
				current = d // это общий делитель и нужно обновить
			}
			bigD := a / d    // нужно вычислить соответствующий большой делитель
			if b%bigD == 0 { // если он является и делителем b
				// т.к. d перебираем в порядке возрастания, то bigD перебирается в порядке убывания,
				// поэтому как только нашли - можем вернуть
				return bigD
			}
		}
	}

	return current
}

// EvklidGCD реализация алгоритма Евклида, где делим каждый из множителей на друг друга пока не дойдем до 0
func EvklidGCD(a, b int) int {
	if a <= 0 || b <= 0 {
		return 0
	}

	for a > 0 && b > 0 { // пока не дошли до нулевых остатков от деления
		if a > b {
			a = a % b // если a > b, то а будет остатком от деления a на b
		} else {
			b = b % a // если b > a, то b будет остатком от деления b на a
		}
	}

	if a == 0 { // если по итогу a нулевое (остаток от деления)
		return b // то НОД будет равен b
	}

	return a // иначе НОД будет равен a (если в b остаток от деления == 0)
}

// EvklidWithSwapGCD реализация алгоритма Евклида, с упрощением кода, свапаем элементы, чтобы не делать проверок
func EvklidWithSwapGCD(a, b int) int {
	if a <= 0 || b <= 0 {
		return 0
	}

	for b > 0 { // пока не дошли до нулевых остатков от деления
		a = a % b   // получаем остаток от деления
		a, b = b, a // меняем числа местами, чтобы при следующей итерации делать ту же самую операцию
	}

	return a // возвращаем a
}

// EvklidRecursiveGCD реализация алгоритма Евклида, с рекурсией
func EvklidRecursiveGCD(a, b int) int {
	if a <= 0 {
		return 0
	}

	if b > 0 { // пока не дошли до нулевых остатков от деления
		return EvklidRecursiveGCD(b, a%b) // рекурсивный вызов с заменой местами + сразу тут же берем остаток от деления
	}

	return a // возвращаем a
}

func main() {
	a, b := 40000000, 530000000

	// запуск для медленно алгоритма
	startTimeForBad := time.Now()
	result := BadAlgoForGCD(a, b)
	timing1 := time.Since(startTimeForBad)
	fmt.Printf("GCD for %d & %d: %d\n", a, b, result)
	fmt.Printf("Time: %0.fm %0.fs %dms %dns\n",
		timing1.Minutes(), timing1.Seconds(), timing1.Microseconds(), timing1.Nanoseconds())

	// запуск ускоренного варианта с делителем
	startTimeForMoreEffective := time.Now()
	result = MoreEffectiveGCD(a, b)
	timing2 := time.Since(startTimeForMoreEffective)
	fmt.Printf("GCD for %d & %d: %d\n", a, b, result)
	fmt.Printf("Time: %0.fm %0.fs %dms %dns\n",
		timing2.Minutes(), timing2.Seconds(), timing2.Microseconds(), timing2.Nanoseconds())

	// запуск варианта с алгоритмом Евклида
	startTimeForEvklid := time.Now()
	result = EvklidGCD(a, b)
	timing3 := time.Since(startTimeForEvklid)
	fmt.Printf("GCD for %d & %d: %d\n", a, b, result)
	fmt.Printf("Time: %0.fm %0.fs %dms %dns\n",
		timing3.Minutes(), timing3.Seconds(), timing3.Microseconds(), timing3.Nanoseconds())

	// запуск варианта с алгоритмом Евклида со swap'ом
	startTimeForEvklidSwap := time.Now()
	result = EvklidWithSwapGCD(a, b)
	timing4 := time.Since(startTimeForEvklidSwap)
	fmt.Printf("GCD for %d & %d: %d\n", a, b, result)
	fmt.Printf("Time: %0.fm %0.fs %dms %dns\n",
		timing4.Minutes(), timing4.Seconds(), timing4.Microseconds(), timing4.Nanoseconds())

	// запуск варианта с алгоритмом Евклида с рекурсией
	startTimeForEvklidRecursive := time.Now()
	result = EvklidRecursiveGCD(a, b)
	timing5 := time.Since(startTimeForEvklidRecursive)
	fmt.Printf("GCD for %d & %d: %d\n", a, b, result)
	fmt.Printf("Time: %0.fm %0.fs %dms %dns\n",
		timing5.Minutes(), timing5.Seconds(), timing5.Microseconds(), timing5.Nanoseconds())
}
