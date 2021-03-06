package main

import (
	"bufio"
	"fmt"
	"gb-go-step-1/homework10/homework4/sort"
	"os"
	"strconv"
)

func main() {
	inputNums := []int{}
	scanner := bufio.NewScanner(os.Stdin)

	println("Enter numbers: ")

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)

		if err != nil {
			panic(err)
		}

		inputNums = append(inputNums, int(num))
	}

	sort.SortByInserts(inputNums)

	println("Sorted list: ")

	for i := range inputNums {
		fmt.Println(inputNums[i])
	}
}
