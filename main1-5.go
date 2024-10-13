package main

import (
	"fmt"
)

func findSubstring(str, substr string, index int) int {
	if len(str) < len(substr) {
		return -1
	}
	if str[:len(substr)] == substr {
		return index
	}
	return findSubstring(str[1:], substr, index+1)
}

func FindFirstPosition(str, substr string) int {
	return findSubstring(str, substr, 0)
}

func main() {
	var str, substr string

	fmt.Println("Введите строку:")
	fmt.Scanln(&str)

	fmt.Println("Введите подстроку:")
	fmt.Scanln(&substr)

	result := FindFirstPosition(str, substr)

	if result != -1 {
		fmt.Printf("Подстрока найдена на позиции: %d\n", result)
	} else {
		fmt.Println("Подстрока не найдена")
	}
}
