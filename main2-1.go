package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var calculator_numbers [2]int
	var calculation string
	fmt.Println("Введите первое число: ")
	fmt.Fscan(os.Stdin, &calculator_numbers[0])
	fmt.Println("Введите второе число: ")
	fmt.Fscan(os.Stdin, &calculator_numbers[1])
	fmt.Println("Введите математическую операцию: ")
	fmt.Fscan(os.Stdin, &calculation)
	switch calculation {
	case "+":
		fmt.Println("Ответ: ", calculator_numbers[0]+calculator_numbers[1])
	case "-":
		fmt.Println("Ответ: ", calculator_numbers[0]+calculator_numbers[1])
	case "*":
		fmt.Println("Ответ: ", calculator_numbers[0]*calculator_numbers[1])
	case "/":
		if calculator_numbers[1] == 0 {
			fmt.Println("Ошибка: На ноль делить нельзя!")
		} else {
			fmt.Println("Ответ: ", calculator_numbers[0]/calculator_numbers[1])
		}
	case "^":
		fmt.Println("Ответ: ", math.Pow(float64(calculator_numbers[0]), float64(calculator_numbers[1])))
	case "%":
		fmt.Println("Ответ: ", calculator_numbers[0]%calculator_numbers[1])
	default:
		fmt.Println("Ошибка: Недопустимая операция!")
	}

}
