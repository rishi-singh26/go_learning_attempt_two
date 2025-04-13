package functionsarevalues

import "fmt"

type transformFM func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 6, 7, 8}

	doubledNumbers := modifyNumbers(&numbers, getTransformFn(&numbers))

	fmt.Println(doubledNumbers)

	tripledNumbers := modifyNumbers(&numbers, getTransformFn(&moreNumbers))

	fmt.Println(tripledNumbers)
}

func modifyNumbers(numbers *[]int, modifier transformFM) (result []int) {
	for _, val := range *numbers {
		result = append(result, modifier(val))
	}
	return
}

func getTransformFn(numbers *[]int) transformFM {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}
