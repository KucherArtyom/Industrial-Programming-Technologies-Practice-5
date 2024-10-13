# Industrial-Programming-Technologies-Practice-5 (ЭФМО-02-24 Кучер Артем Сергеевич)
## Задачи на языке программирования Go
### 1. Линейное программирование

**1) Перевод чисел из одной системы счисления в другую**
```
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
```
**2) Решение квадратного уравнения**
```
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Print("Введите коэффициент a: ")
	fmt.Scan(&a)
	fmt.Print("Введите коэффициент b: ")
	fmt.Scan(&b)
	fmt.Print("Введите коэффициент c: ")
	fmt.Scan(&c)
	solveQuadraticEquation(a, b, c)
}

func solveQuadraticEquation(a, b, c float64) {
	d := b*b - 4*a*c

	if d > 0 {
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		x2 := (-b - math.Sqrt(d)) / (2 * a)
		fmt.Printf("Уравнение имеет два действительных корня: x1 = %.2f, x2 = %.2f\n", x1, x2)
	} else if d == 0 {
		x := -b / (2 * a)
		fmt.Printf("Уравнение имеет один действительный корень: x = %.2f\n", x)
	} else {
		realPart := -b / (2 * a)
		imaginaryPart := math.Sqrt(-d) / (2 * a)
		fmt.Printf("Уравнение имеет два комплексных корня: x1 = %.2f + %.2fi, x2 = %.2f - %.2fi\n", realPart, imaginaryPart, realPart, imaginaryPart)
	}
}
```
**3) Сортировка чисел по модулю**
```
package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func sortByAbs(numbers []int) []int {
	sort.Slice(numbers, func(i, j int) bool {
		return math.Abs(float64(numbers[i])) < math.Abs(float64(numbers[j]))
	})
	return numbers
}

func main() {
	numbers := make([]int, 3)
	fmt.Println("Введите первый элемент трёхмерного массива")
	fmt.Fscan(os.Stdin, &numbers[0])
	fmt.Println("Введите второй элемент трёхмерного массива")
	fmt.Fscan(os.Stdin, &numbers[1])
	fmt.Println("Введите третий элемент трёхмерного массива")
	fmt.Fscan(os.Stdin, &numbers[2])
	fmt.Println("Отсортированный массив по модулю:", sortByAbs(numbers))
}
```
**4) Слияние двух отсортированных массивов**
```
package main

import (
	"fmt"
)

func mergeSortedArrays(arr1, arr2 []int) []int {
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}

	if arr1[0] < arr2[0] {
		return append([]int{arr1[0]}, mergeSortedArrays(arr1[1:], arr2)...)
	}
	return append([]int{arr2[0]}, mergeSortedArrays(arr1, arr2[1:])...)
}

func inputArray() []int {
	var n int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scanln(&n)

	arr := make([]int, n)
	fmt.Println("Введите элементы массива по порядку:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	return arr
}

func main() {
	fmt.Println("Введите первый отсортированный массив:")
	arr1 := inputArray()
	fmt.Println("Введите второй отсортированный массив:")
	arr2 := inputArray()
	result := mergeSortedArrays(arr1, arr2)
	fmt.Println("Результат слияния:", result)
}
```
**5) Нахождение подстроки в строке без использования встроенных функций**
```
package main

import (
	"fmt"
)

func findSubstring(str, substr string, index int) int {
	if len(str) < len(substr) {
		return -1
	}
	if str[:len(substr)] == substr {
		return index
	}
	return findSubstring(str[1:], substr, index+1)
}

func FindFirstPosition(str, substr string) int {
	return findSubstring(str, substr, 0)
}

func main() {
	var str, substr string

	fmt.Println("Введите строку:")
	fmt.Scanln(&str)

	fmt.Println("Введите подстроку:")
	fmt.Scanln(&substr)

	result := FindFirstPosition(str, substr)

	if result != -1 {
		fmt.Printf("Подстрока найдена на позиции: %d\n", result)
	} else {
		fmt.Println("Подстрока не найдена")
	}
}
```
### 2. Условия
**1) Калькулятор с расширенными операциями**
```
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
```
**2) Проверка палиндрома**
```
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	var cleanedStr []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleanedStr = append(cleanedStr, r)
		}
	}

	for i := 0; i < len(cleanedStr)/2; i++ {
		if cleanedStr[i] != cleanedStr[len(cleanedStr)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("Введите строку:")
	var input string
	fmt.Scanln(&input)
	if isPalindrome(input) {
		fmt.Println("Строка является палиндромом.")
	} else {
		fmt.Println("Строка не является палиндромом.")
	}
}
```
**3) Нахождение пересечения трех отрезков**
```
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
```
**4) Выбор самого длинного слова в предложении**
```
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
```
**5) Проверка высокосного года**
```
package main

import (
	"fmt"
	"os"
)

func main() {
	var year int
	fmt.Println("Является ли данный год вискокосным? Введите год: ")
	fmt.Fscan(os.Stdin, &year)

	if year%400 == 0 {
		fmt.Println("True")
	} else if year%100 == 0 {
		fmt.Println("False")
	} else if year%4 == 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
```
### 3. Циклы
**1) Числа Фибоначчи до определенного числа**
```
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
```
**2) Определение простых чисел в диапазоне**
```
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
```
**3) Числа Армстронга в заданном диапазоне**
```
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
```
**4) Реверс строки**
```
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
```
**5) Нахождение наибольшего общего делителя (НОД)**
```
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
```
