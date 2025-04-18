package main

import (
	"fmt"

	"rishisingh.in/practice-project/filemanager"
	"rishisingh.in/practice-project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errChans[index] = make(chan error)
		fileManager := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdManager := cmdmanager.New()

		priceJob := prices.NewTaxIncludedPriceJob(fileManager, taxRate)
		go priceJob.Process(doneChans[index], errChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// 	return
		// }
	}

	for index := range taxRates {
		select {
		case err := <-errChans[index]:
			if err != nil {
				fmt.Println(err)
				return
			}
		case <-doneChans[index]:
			fmt.Println("Done")
		}
	}
}
