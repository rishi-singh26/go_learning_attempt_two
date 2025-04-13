package lists

import "fmt"

func main() {
	prices := []float64{12.2, 123.2}

	fmt.Println(prices[1])

	discountedPrices := []float64{121.2, 12.2, 34.99}

	mergedPrices := append(prices, discountedPrices...)
	fmt.Println(mergedPrices)
}

// func main() {
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	productNames := [4]string{"A Book"}

// 	productNames[2] = "A Carpet"

// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:3] // Create slice method 1
// 	fmt.Println(featuredPrices)

// 	featuredPrices = prices[:3] // Create slice method 2
// 	fmt.Println(featuredPrices)

// 	featuredPrices = prices[1:] // Create slice method 3
// 	fmt.Println(featuredPrices)

// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)
// }
