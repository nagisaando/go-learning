package main

import (
	"fmt"

	"example.com/price-calculator-2/filemanager"
	"example.com/price-calculator-2/price"
)

func main() {

	taxRates := []float64{0, 10, 20}
	doneChans := make([]chan bool, len(taxRates))

	// Function with go keyword won't be able to return value.
	// If we want to return value, we have to use channel
	errorChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)
		// we can interchange struct because of IOManager interface. whether we use fileManager or CMDLineManager to price struct, it won't show error
		fileManager := filemanager.New("price.txt", fmt.Sprintf("result_%v.json", taxRate))
		// CMDLineManager := cmdmanager.New()

		priceJob := price.NewTaxIncludedPriceJob(fileManager, taxRate)

		go priceJob.Process(doneChans[i], errorChans[i])
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}

	// "select" wait for the different values to be emitted through channel.
	// it works when we want to get either channel value (doneChan or errorChan in this case)
	// it will be useful when we just want to get one value from multiple channels

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:

			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("done!")
		}
	}
	// for _, doneChan := range doneChans {
	// 	<-doneChan // tells GO that we wait until every channel has emitted one value
	// }

}
