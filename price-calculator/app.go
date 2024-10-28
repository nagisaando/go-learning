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
			priceWithTax := tax.New(currentPriceData)
			fmt.Println(priceWithTax)
			priceWithTax.Save()

		}
	}

}

func askOptions() (string, []float64, error) {
	fmt.Println("=======================================")
	fmt.Printf("Welcome, we have the saved data: \n\n")

	fmt.Printf("Price: \n")
	prices, _ := price.Read()
	fmt.Println(prices)

	fmt.Printf("Price with tax: \n")

	pricesWithTax, _ := tax.Read()
	fmt.Println(pricesWithTax)

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
