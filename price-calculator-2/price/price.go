package price

import (
	"fmt"

	"example.com/price-calculator-2/conversion"
	"example.com/price-calculator-2/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"` // adding minus will exclude the key from json
}

func NewTaxIncludedPriceJob(ioManager iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{0, 10, 20},
		TaxRate:     taxRate,
		IOManager:   ioManager,
		// we don't initialize taxIncludedPrices at the moment since it needs to be calculated later
	}
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadData()
	if err != nil {
		errorChan <- err
		return
	}
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		// fmt.Sprintf can return float64 to string

		taxIncludedPrice := price * (1 + (job.TaxRate / 100))
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	err = job.IOManager.WriteResult(job)
	if err != nil {
		errorChan <- err
		return
	}

	doneChan <- true

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

	// this update won't work if we don't get job data as pointer
	// works: func (job *TaxIncludedPriceJob)
	// does NOT work: func (job TaxIncludedPriceJob) since it is just a copy and can not update the original data

	job.InputPrices = prices
	return nil

}
