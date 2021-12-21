package main

import (
	"fmt"
	"gb-go-step-1/homework10/homework5/fibb"
)

func main() {
	var number int
	fmt.Println("Введите число: ")
	fmt.Scan(&number)

	fmt.Printf("Число Фибоначчи: %v\n", fibb.GetFibonacciByCache(number))
}
