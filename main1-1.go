package main

import (
	"fmt"
	"os"
)

func decimalToBinary(n int) string {
	if n == 0 {
		return "0"
	}
	binary := ""
	for n > 0 {
		bit := n % 2
		binary = fmt.Sprintf("%d%s", bit, binary)
		n /= 2
	}
	return binary
}

func main() {
	var number10 int
	fmt.Println("Введите число в десятичной системе счисления: ")
	_, err := fmt.Fscan(os.Stdin, &number10)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	binaryOutput := decimalToBinary(number10)
	fmt.Printf("Число в двоичной системе счисления: %s\n", binaryOutput)
}
