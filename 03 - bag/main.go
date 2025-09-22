package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type Item struct {
	weight float64
	value  float64
}

// getProfit получение стоимости веса вещи
func (i *Item) getProfit() float64 {
	return i.weight / i.value
}

// generateItems генерирует случайный набор данных
func generateItems(count int, seed int64) []Item {
	rand.Seed(seed)
	items := make([]Item, count)

	for i := 0; i < count; i++ {
		// Генерация веса от 0.1 до 10 кг
		weight := 0.1 + rand.Float64()*9.9
		// Генерация ценности от 1 до 100 условных единиц
		value := 1.0 + rand.Float64()*99.0

		items[i] = Item{
			weight: math.Ceil(weight),
			value:  math.Ceil(value),
		}
	}
	return items
}

// GetMaxKnapsackValue наивный вариант
func GetMaxKnapsackValue(capacity float64, items []Item) (value float64) {
	fmt.Printf("Несортированный список: %v\n", items)

	// пока в рюкзаке есть место
	for _, item := range items {
		if capacity > item.weight {
			// можем взять вещь целиком и продолжить
			capacity -= item.weight
			value += item.value
		} else {
			value += item.value * (capacity / item.weight)
			break
		}
	}

	return value
}

// GetMaxKnapsackValueFromSorted вариант с отсортированными по качеству вещами
func GetMaxKnapsackValueFromSorted(capacity float64, items []Item) (value float64) {
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].getProfit() < items[j].getProfit()
	})

	fmt.Printf("Отсортированный список: %v\n", items)

	// пока в рюкзаке есть место
	for _, item := range items {
		if capacity > item.weight {
			// можем взять вещь целиком и продолжить
			capacity -= item.weight
			value += item.value
		} else {
			value += item.value * (capacity / item.weight)
			break
		}
	}

	return value
}

// GetMaxKnapsackValueWoutDelim вариант с отсортированными по качеству вещами, без деления
func GetMaxKnapsackValueWoutDelim(capacity float64, items []Item) (value float64) {
	// деление в правой и левой частях заменяем на умножение
	// l.w / l.v < r.w / r.v ===>>> l.w * r.v < r.w * l.v
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].weight*items[j].value < items[j].weight*items[i].value
	})

	fmt.Printf("Отсортированный список: %v\n", items)

	// пока в рюкзаке есть место
	for _, item := range items {
		if capacity > item.weight {
			// можем взять вещь целиком и продолжить
			capacity -= item.weight
			value += item.value
		} else {
			value += item.value * (capacity / item.weight)
			break
		}
	}

	return value
}

func main() {
	// Фиксированный seed для повторяемости результатов
	seed := int64(42)
	items := generateItems(100, seed)

	capacity := float64(3)

	result1 := GetMaxKnapsackValue(capacity, items)
	fmt.Printf("Результат работы наивного алгоритма: %.0f\n", result1)

	result2 := GetMaxKnapsackValueFromSorted(capacity, items)
	fmt.Printf("Результат работы с сортировкой: %.0f\n", result2)

	result3 := GetMaxKnapsackValueWoutDelim(capacity, items)
	fmt.Printf("Результат работы с сортировкой, без деления: %.0f\n", result3)
}
