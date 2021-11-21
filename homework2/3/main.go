package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Print("Введите трёхзначное число: ")
	fmt.Scanln(&number)

	hundredsAmount := number / 100
	tensAmount := number / 10 % 10
	unitsAmount := number % 10

	fmt.Printf("Количество сотен: %v\nКоличество десятков: %v\nКоличество единиц: %v\n", hundredsAmount, tensAmount, unitsAmount)
}
