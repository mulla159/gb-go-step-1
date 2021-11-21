package main

import "fmt"

func main() {
	var aLength, bLength int

	fmt.Println("Введите длину строны a: ")
	fmt.Scan(&aLength)

	fmt.Println("Введите длину строны b: ")
	fmt.Scan(&bLength)

	fmt.Println("Площадь прямоугольника: ", aLength*bLength)
}
