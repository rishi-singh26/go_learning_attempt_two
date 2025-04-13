package prices

import (
	"fmt"

	"rishisingh.in/practice-project/conversion"
	"rishisingh.in/practice-project/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"taxRate"`
	InputPrices       []float64           `json:"inputPrices"`
	TaxIncludedPrices map[string]string   `json:"taxIncludedPrices"`
	IOManager         iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return err
	}

	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxincludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxincludedPrice)
	}

	job.TaxIncludedPrices = result
	fmt.Println(result)
	job.IOManager.WriteResult(job)

	return nil
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   iom,
	}
}
