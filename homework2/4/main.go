package main

import (
	"fmt"
)

func main() {
	var inputNumber int

	fmt.Print("Введите целое число: ")
	fmt.Scanln(&inputNumber)

	fmt.Printf("Простые числа от 0 до %v:\n", inputNumber)

	// перебираем все числа от 2 до введенного и выводим простые
	for i := 2; i <= inputNumber; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

// функция проверяет является ли число простым
func isPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return number > 1
}
