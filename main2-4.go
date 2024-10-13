package main

import (
	"fmt"
)

func reverseString(str string) string {
	if len(str) == 0 {
		return str
	}
	return string(str[len(str)-1]) + reverseString(str[:len(str)-1])
}

func main() {
	var input string
	fmt.Println("Введите строку:")
	fmt.Scanln(&input)

	reversed := reverseString(input)
	fmt.Println("Строка в обратном порядке:", reversed)
}
