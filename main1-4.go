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
