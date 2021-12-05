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
	fibNumbers := map[int]int{0: 0, 1: 1}

	if val, exists := fibNumbers[n]; exists {
		return val
	}

	fibNumbers[n] = getFibonacci(n-1) + getFibonacci(n-2)
	return fibNumbers[n]
}
