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
