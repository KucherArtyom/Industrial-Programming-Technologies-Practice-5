package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for {
		b = a % b
		a, b = b, a

		if b == 0 {
			return a
		}
	}
}

func main() {
	var a, b int

	fmt.Println("Введите два числа:")
	fmt.Scan(&a, &b)

	result := gcd(a, b)
	fmt.Printf("Наибольший общий делитель: %d\n", result)
}
