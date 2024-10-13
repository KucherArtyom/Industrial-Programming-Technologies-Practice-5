package main

import (
	"fmt"
	"math"
)

func hasIntersection(x1, y1, x2, y2, x3, y3 int) bool {
	start := int(math.Max(float64(x1), math.Max(float64(x2), float64(x3))))
	end := int(math.Min(float64(y1), math.Min(float64(y2), float64(y3))))
	return start <= end
}

func main() {
	var x1, y1, x2, y2, x3, y3 int
	fmt.Println("Введите начальную и конечную точки первого отрезка:")
	fmt.Scan(&x1, &y1)
	fmt.Println("Введите начальную и конечную точки второго отрезка:")
	fmt.Scan(&x2, &y2)
	fmt.Println("Введите начальную и конечную точки третьего отрезка:")
	fmt.Scan(&x3, &y3)

	if hasIntersection(x1, y1, x2, y2, x3, y3) {
		fmt.Println("Область пересечения существует.")
	} else {
		fmt.Println("Область пересечения не существует.")
	}
}
