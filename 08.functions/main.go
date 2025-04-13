package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	doubledNumbers := modifyNumbers(&numbers, double)

	fmt.Println(doubledNumbers)

	tripledNumbers := modifyNumbers(&numbers, triple)

	fmt.Println(tripledNumbers)
}

func modifyNumbers(numbers *[]int, modifier func(int) int) (result []int) {
	for _, val := range *numbers {
		result = append(result, modifier(val))
	}
	return
}

func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}
