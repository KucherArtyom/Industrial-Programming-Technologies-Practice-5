package main

import (
	"fmt"
	"math"
)

func main() {
	// Ввод коэффициентов
	var a, b, c float64

	fmt.Print("Введите коэффициент a: ")
	fmt.Scan(&a)
	fmt.Print("Введите коэффициент b: ")
	fmt.Scan(&b)
	fmt.Print("Введите коэффициент c: ")
	fmt.Scan(&c)

	// Решение уравнения
	solveQuadraticEquation(a, b, c)
}

func solveQuadraticEquation(a, b, c float64) {
	// Вычисляем дискриминант
	d := b*b - 4*a*c

	if d > 0 {
		// Два действительных корня
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		x2 := (-b - math.Sqrt(d)) / (2 * a)
		fmt.Printf("Уравнение имеет два действительных корня: x1 = %.2f, x2 = %.2f\n", x1, x2)
	} else if d == 0 {
		// Один действительный корень
		x := -b / (2 * a)
		fmt.Printf("Уравнение имеет один действительный корень: x = %.2f\n", x)
	} else {
		// Комплексные корни
		realPart := -b / (2 * a)
		imaginaryPart := math.Sqrt(-d) / (2 * a)
		fmt.Printf("Уравнение имеет два комплексных корня: x1 = %.2f + %.2fi, x2 = %.2f - %.2fi\n", realPart, imaginaryPart, realPart, imaginaryPart)
	}
}
