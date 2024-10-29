package main

import (
	"fmt"

	"example.com/price-calculator-2/filemanager"
	"example.com/price-calculator-2/price"
)

func main() {

	taxRates := []float64{0, 10, 20}

	for _, taxRate := range taxRates {
		fileManager := filemanager.New("price.txt", fmt.Sprintf("result_%v.json", taxRate))
		priceJob := price.NewTaxIncludedPriceJob(*fileManager, taxRate)

		priceJob.Process()
	}

}
