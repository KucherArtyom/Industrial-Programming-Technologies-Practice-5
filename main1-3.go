package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func sortByAbs(numbers []int) []int {
	// Используем встроенную сортировку
	sort.Slice(numbers, func(i, j int) bool {
		// Сравниваем по модулю
		return math.Abs(float64(numbers[i])) < math.Abs(float64(numbers[j]))
	})
	return numbers
}

func main() {
	// Инициализация слайса на 3 элемента
	numbers := make([]int, 3)

	fmt.Println("Введите первый элемент трёхмерного массива")
	fmt.Fscan(os.Stdin, &numbers[0])
	fmt.Println("Введите второй элемент трёхмерного массива")
	fmt.Fscan(os.Stdin, &numbers[1])
	fmt.Println("Введите третий элемент трёхмерного массива")
	fmt.Fscan(os.Stdin, &numbers[2])

	// Вывод отсортированного массива
	fmt.Println("Отсортированный массив по модулю:", sortByAbs(numbers))
}
