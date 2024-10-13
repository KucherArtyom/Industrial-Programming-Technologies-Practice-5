package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var start, end int
	fmt.Println("Введите начальное число диапазона: ")
	fmt.Fscan(os.Stdin, &start)
	fmt.Println("Введите конечное число диапазона: ")
	fmt.Fscan(os.Stdin, &end)

	fmt.Println("Числа Армстронга в заданном диапазоне:")
	findArmstrongNumbers(start, end)
}

func findArmstrongNumbers(start, end int) {
	for i := start; i <= end; i++ {
		if isArmstrong(i) {
			fmt.Println(i)
		}
	}
}

func isArmstrong(num int) bool {
	numStr := fmt.Sprint(num)
	n := len(numStr)

	sum := 0
	temp := num
	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(n)))
		temp /= 10
	}

	return sum == num
}
