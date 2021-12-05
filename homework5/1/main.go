package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Введите число: ")
	fmt.Scan(&number)

	fmt.Printf("Число Фибоначчи: %v\n", getFibonacci(number))
}

func getFibonacci(n int) int {
	var result int

	if n <= 1 {
		result = n
	} else {
		result = getFibonacci(n-1) + getFibonacci(n-2)
	}

	return result
}
