package main

import (
	"fmt"

	"example.com/price-calculator-2/cmdmanager"
	"example.com/price-calculator-2/price"
)

func main() {

	taxRates := []float64{0, 10, 20}

	for _, taxRate := range taxRates {
		// we can interchange struct because of IOManager interface. whether we use fileManager or CMDLineManager to price struct, it won't show error
		// fileManager := filemanager.New("price.txt", fmt.Sprintf("result_%v.json", taxRate))
		CMDLineManager := cmdmanager.New()

		priceJob := price.NewTaxIncludedPriceJob(CMDLineManager, taxRate)

		err := priceJob.Process()
		if err != nil {
			fmt.Println(err)
		}
	}

}
