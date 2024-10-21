package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	newPriceWithTaxRate := map[int][]float64{}

	var taxRates = []int{0, 10, 20}

	var prices = scanPrice()

	fmt.Print(prices)

	// get input for prices and store to text file

	// take out the price from the text

	// calculate new price based on tax
	for val, _ := range taxRates {
		newPriceWithTaxRate[val] = append(newPriceWithTaxRate[val], 10)
	}

	// print result
	fmt.Println(newPriceWithTaxRate)

	// store data in json format
	// {0: [10, 20, 30], 10: [11, 22, 33]}

}

func scanPrice() []float64 {

	fmt.Println("Please enter prices separated by space")

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return []float64{}
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // for windows since return key return as \r\n

	stringPrices := strings.Split(text, " ")
	numberPrices := []float64{}

	for i, val := range stringPrices {
		convertedPrice, _ := strconv.ParseFloat(val, 64)
		numberPrices[i] = convertedPrice

	}

	return numberPrices

}
