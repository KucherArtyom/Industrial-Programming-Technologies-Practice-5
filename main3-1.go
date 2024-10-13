package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	var number1 = 0
	var number2 = 1
	var number3 = 0
	fmt.Println("Введите любое целое число n: ")
	fmt.Fscan(os.Stdin, &n)
	for i := 1; i <= n; i++ {
		fmt.Println(number1)
		number3 = number1 + number2
		number1 = number2
		number2 = number3
	}
}
