package main

import (
	"fmt"
	"gb-go-step-1/homework10/homework3/factorial"
	"math"
	"os"
)

func main() {
	var a float64
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	var operation string
	fmt.Print("Введите операцию (+, -, *, /, ^, !): ")
	fmt.Scanln(&operation)

	var b float64
	if operation != "!" {
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
	}

	var result float64
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Делить на ноль нельзя!")
			os.Exit(1)
		}

		result = a / b
	case "^":
		result = math.Pow(a, b)
	case "!":
		result = factorial.Factorial(a)
	default:
		fmt.Println("Некорректная операция!")
		os.Exit(1)
	}

	fmt.Print("Результат выполнения операции: ")

	if isInt(result) {
		fmt.Printf("%v\n", result)
	} else {
		fmt.Printf("%.2f\n", result)
	}
}

func isInt(number float64) bool {
	return number == float64(int(number))
}
