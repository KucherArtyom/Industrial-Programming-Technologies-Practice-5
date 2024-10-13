package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Функция для проверки палиндрома
func isPalindrome(s string) bool {
	// Приводим строку к нижнему регистру
	s = strings.ToLower(s)

	// Создаем новую строку только из букв и цифр, игнорируя пробелы и знаки препинания
	var cleanedStr []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleanedStr = append(cleanedStr, r)
		}
	}

	// Проверка, является ли строка палиндромом
	for i := 0; i < len(cleanedStr)/2; i++ {
		if cleanedStr[i] != cleanedStr[len(cleanedStr)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	// Пример ввода
	fmt.Println("Введите строку:")
	var input string
	fmt.Scanln(&input)

	// Вывод результата
	if isPalindrome(input) {
		fmt.Println("Строка является палиндромом.")
	} else {
		fmt.Println("Строка не является палиндромом.")
	}
}
