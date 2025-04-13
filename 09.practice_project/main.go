package main

import (
	"fmt"

	"rishisingh.in/practice-project/cmdmanager"
	"rishisingh.in/practice-project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fileManager := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdManager := cmdmanager.New()
		err := prices.NewTaxIncludedPriceJob(cmdManager, taxRate).Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
			return
		}
	}
}
