package main

import (
	"fmt"
	"math"
)

func main() {
	var circleArea, circleDiameter, circleRadius, circleLength float64

	fmt.Print("Введите площадь круга: ")
	fmt.Scanln(&circleArea)

	circleRadius = math.Sqrt(circleArea / math.Pi)

	circleDiameter = circleRadius * 2

	circleLength = 2 * math.Pi * circleRadius

	fmt.Printf("Диаметр окружности: %v\nДлина окружностика: %v\n", circleDiameter, circleLength)
}
