package prices

import (
	"fmt"

	"rishisingh.in/practice-project/conversion"
	"rishisingh.in/practice-project/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64                 `json:"taxRate"`
	InputPrices       []float64               `json:"inputPrices"`
	TaxIncludedPrices map[string]string       `json:"taxIncludedPrices"`
	IOManager         filemanager.FileManager `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxincludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxincludedPrice)
	}

	job.TaxIncludedPrices = result
	fmt.Println(result)
	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   fm,
	}
}
