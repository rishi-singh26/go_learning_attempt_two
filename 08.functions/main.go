package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sumup(numbers...))
	fmt.Println(sumup(1, 2, 3, 4, 5, 7, -2))
}

func sumup(numbers ...int) int {
	result := 0

	for _, val := range numbers {
		result += val
	}

	return result
}
