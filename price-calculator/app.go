package main

import (
	"errors"
	"fmt"

	"example.com/price-calculator/price"
	"example.com/price-calculator/scan"
	"example.com/price-calculator/tax"
)

func main() {

	for {
		option, currentPriceData, err := askOptions()

		if err != nil {
			fmt.Print(err)
			return
		}

		switch option {
		case "3":

			return
		case "1":
			prices, err := price.Scan()
			fmt.Println(prices)
			if err != nil {
				fmt.Println(err)
				continue
			}

			price.Save(prices)
		case "2":

			// newPriceWithTaxRate := map[int][]float64{}

			// // calculate new price based on tax
			// for _, tax := range taxRates {
			// 	for _, price := range currentPriceData {

			// 		newPrice := price + (price * (float64(tax) / 100.00))

			// 		newPriceWithTaxRate[tax] = append(newPriceWithTaxRate[tax], newPrice)

			// 	}

			// }
			priceWithTax := tax.New(currentPriceData)
			// print result
			fmt.Println(*priceWithTax)

			// store data in json format
			// {0: [10, 20, 30], 10: [11, 22, 33]}

		}
	}

}

func askOptions() (string, []float64, error) {
	fmt.Println("=======================================")
	fmt.Printf("Welcome, we have the saved prices data: \n")

	prices, _ := price.Read()
	fmt.Println(prices)

	fmt.Printf("What do you need? Please enter number: \n\n")
	fmt.Println("1. Scan prices")
	fmt.Println("2. Calculate with tax")
	fmt.Println("3. Exit")
	fmt.Println("=======================================")

	option, err := scan.ReadTextFromCML()

	if err != nil || (option != "1" && option != "2" && option != "3") {
		return "", []float64{}, errors.New("Please enter valid option")
	}
	return option, prices, err
}
