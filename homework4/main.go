package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputNums := []int64{}
	scanner := bufio.NewScanner(os.Stdin)

	println("Enter numbers: ")

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)

		if err != nil {
			panic(err)
		}

		inputNums = append(inputNums, num)
	}

	sortArrayInAsc(inputNums)

	println("Sorted list: ")

	for i := range inputNums {
		fmt.Println(inputNums[i])
	}
}

func sortArrayInAsc(arr []int64) {
	for i := 1; i < len(arr); i++ {
		currentIndex := i

		for j := i - 1; j >= 0; j-- {
			if arr[currentIndex] < arr[j] {
				arr[currentIndex], arr[j] = arr[j], arr[currentIndex]

				currentIndex = j
			} else {
				continue
			}
		}
	}
}
