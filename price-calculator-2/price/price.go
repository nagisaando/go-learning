package price

import (
	"fmt"

	"example.com/price-calculator-2/conversion"
	"example.com/price-calculator-2/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{0, 10, 20},
		TaxRate:     taxRate,
		// we don't initialize taxIncludedPrices at the moment since it needs to be calculated later
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		// fmt.Sprintf can return float64 to string

		taxIncludedPrice := price * (1 + (job.TaxRate / 100))
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	filemanager.WriteJSON(fmt.Sprintf("result_%v.json", job.TaxRate), job)

}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := filemanager.ReadLines("price.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return

	}

	// this update won't work if we don't get job data as pointer
	// works: func (job *TaxIncludedPriceJob)
	// does NOT work: func (job TaxIncludedPriceJob) since it is just a copy and can not update the original data

	job.InputPrices = prices

}
