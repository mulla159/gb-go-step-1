package fibb

var fibNumbers = map[int]int{0: 0, 1: 1}

func GetFibonacciByCache(n int) int {
	if val, exists := fibNumbers[n]; exists {
		return val
	}

	fibNumbers[n] = GetFibonacciByCache(n-1) + GetFibonacciByCache(n-2)
	return fibNumbers[n]
}

func GetFibonacciRecur(n int) int {
	var result int

	if n <= 1 {
		result = n
	} else {
		result = GetFibonacciRecur(n-1) + GetFibonacciRecur(n-2)
	}

	return result
}
