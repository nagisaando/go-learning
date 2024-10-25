package main

import "example.com/price-calculator-2/price"

func main() {

	taxRates := []float64{0, 10, 20}

	for _, taxRate := range taxRates {
		priceJob := price.NewTaxIncludedPriceJob(taxRate)

		priceJob.Process()
	}

}
