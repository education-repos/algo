package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// Node Узел дерева Хаффмана
type Node struct {
	char  rune
	freq  int
	left  *Node
	right *Node
}

// PriorityQueue Реализация интерфейса heap.Interface для приоритетной очереди
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]

	return item
}

// Функция для построения дерева Хаффмана
func buildHuffmanTree(text string) *Node {
	// Подсчет частот символов
	freqMap := make(map[rune]int)
	for _, char := range text {
		freqMap[char]++
	}

	// Создание приоритетной очереди
	pq := make(PriorityQueue, 0, len(freqMap))

	for char, freq := range freqMap {
		pq = append(pq, &Node{char: char, freq: freq})
	}

	heap.Init(&pq)

	// Построение дерева Хаффмана
	for pq.Len() > 1 {
		// Извлекаем два узла с наименьшей частотой
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)

		// Создаем новый узел с суммой частот
		parent := &Node{
			freq:  left.freq + right.freq,
			left:  left,
			right: right,
		}

		// Добавляем новый узел обратно в очередь
		heap.Push(&pq, parent)
	}

	// Возвращаем корень дерева
	return heap.Pop(&pq).(*Node)
}

// Функция для генерации кодов Хаффмана
func generateCodes(root *Node, prefix string, codes map[rune]string) {
	if root == nil {
		return
	}

	// Если это листовой узел (символ)
	if root.char != 0 {
		codes[root.char] = prefix
		return
	}

	// Рекурсивно обходим левое и правое поддеревья
	generateCodes(root.left, prefix+"0", codes)
	generateCodes(root.right, prefix+"1", codes)
}

// Функция кодирования строки
func encode(text string, codes map[rune]string) string {
	encoded := ""
	for _, char := range text {
		encoded += codes[char]
	}
	return encoded
}

// Функция декодирования строки
func decode(encoded string, root *Node) string {
	decoded := ""
	current := root

	for _, bit := range encoded {
		if bit == '0' {
			current = current.left
		} else {
			current = current.right
		}

		// Если достигли листового узла
		if current.left == nil && current.right == nil {
			decoded += string(current.char)
			current = root // Возвращаемся к корню для следующего символа
		}
	}

	return decoded
}

// Вспомогательная функция для красивого вывода кодов
func printCodes(codes map[rune]string) {
	// Сортируем символы для красивого вывода
	var chars []rune
	for char := range codes {
		chars = append(chars, char)
	}

	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	fmt.Println("Коды Хаффмана:")
	for _, char := range chars {
		fmt.Printf("'%c': %s\n", char, codes[char])
	}
	fmt.Println()
}

func main() {
	text := "hello world"

	fmt.Printf("Исходный текст: %s\n\n", text)

	// Построение дерева Хаффмана
	root := buildHuffmanTree(text)

	// Генерация кодов
	codes := make(map[rune]string)
	generateCodes(root, "", codes)

	// Вывод кодов
	printCodes(codes)

	// Кодирование
	encoded := encode(text, codes)
	fmt.Printf("Закодированная строка: %s\n", encoded)
	fmt.Printf("Длина закодированной строки: %d бит\n", len(encoded))
	fmt.Printf("Исходная длина: %d байт (%d бит)\n", len(text), len(text)*8)
	fmt.Printf("Степень сжатия: %.2f%%\n",
		float64(len(text)*8-len(encoded))/float64(len(text)*8)*100)

	// Декодирование
	decoded := decode(encoded, root)
	fmt.Printf("\nДекодированная строка: %s\n", decoded)

	// Проверка корректности
	if text == decoded {
		fmt.Println("✓ Кодирование и декодирование выполнены успешно!")
	} else {
		fmt.Println("✗ Ошибка в кодировании/декодировании!")
	}
}
