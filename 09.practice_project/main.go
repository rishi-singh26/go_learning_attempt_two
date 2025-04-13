package main

import (
	"fmt"

	"rishisingh.in/practice-project/filemanager"
	"rishisingh.in/practice-project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fileManager := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		prices.NewTaxIncludedPriceJob(fileManager, taxRate).Process()
	}
}
