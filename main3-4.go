package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s)
	reversed := make([]rune, len(runes))
	for i := 0; i < len(runes); i++ {
		reversed[i] = runes[len(runes)-1-i]
	}

	return string(reversed)
}

func main() {
	var input string
	fmt.Println("Введите строку:")
	fmt.Scanln(&input)
	reversedOutput := reverseString(input)
	fmt.Printf("Строка в обратном порядке: %s\n", reversedOutput)
}
