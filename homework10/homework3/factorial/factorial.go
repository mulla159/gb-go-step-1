package factorial

func Factorial(number float64) float64 {
	var result float64
	if number == 0 {
		result = 1
	} else {
		result = number * Factorial(number-1)
	}
	return result
}
