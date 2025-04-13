package main

import "fmt"

func main() {
	addInt := add(1, 2)
	addFloat := add(1.0, 2.0)
	addString := add("Name: ", "Rishi")

	fmt.Println(addInt)
	fmt.Println(addFloat)
	fmt.Println(addString)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
