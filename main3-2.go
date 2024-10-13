package main

import (
	"fmt"
	"os"
)

func main() {
	var number1 int
	var number2 int
	fmt.Println("Введите первое целое число:")
	fmt.Fscan(os.Stdin, &number1)
	fmt.Println("Введите второе целое число:")
	fmt.Fscan(os.Stdin, &number2)

	if number2 > 1 {
		fmt.Println("Простые числа от", number1, "до", number2, ":")
		for i := number1; i <= number2; i++ {
			isPrime := true
			for j := 2; j*j <= i; j++ {
				if i%j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				fmt.Print(i, " ")
			}
		}
		fmt.Println()
	} else {
		fmt.Println("Простых чисел нет")
	}
}
