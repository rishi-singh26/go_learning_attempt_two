package recursion

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

func factorial(num int) int {
	if num > 1 {
		return num * factorial(num-1)
	} else {
		return num
	}

	// result := 1

	// for i := 1; i <= num; i++ {
	// 	result *= i
	// }

	// return result
}
