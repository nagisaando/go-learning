package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	newPriceWithTaxRate := map[int][]float64{}

	taxRates := []int{0, 10, 20}

	for {
		option, err := askOptions()

		if err != nil {
			fmt.Print(err)
			break
		}

		switch option {
		case "3":

			return
		case "1":
			prices, err := scanPrice()
			fmt.Println(prices)
			if err != nil {
				fmt.Println(err)
				continue
			}

			savePrices(prices)
		case "2":
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
	}

}

func askOptions() (string, error) {
	fmt.Println("=======================================")
	fmt.Printf("Welcome, we have the saved prices data: \n")

	prices, _ := readPricesData()
	fmt.Println(prices)

	fmt.Printf("What do you need? Please enter number: \n\n")
	fmt.Println("1. Scan prices")
	fmt.Println("2. Calculate with tax")
	fmt.Println("3. Exit")
	fmt.Println("=======================================")

	option, err := readTextFromCML()

	if err != nil || (option != "1" && option != "2" && option != "3") {
		return "", errors.New("Please enter valid option")
	}
	return option, err
}

func trimReturnKey(text *string) {
	*text = strings.TrimSuffix(*text, "\n")
	*text = strings.TrimSuffix(*text, "\r")

}

func readTextFromCML() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	trimReturnKey(&text)
	return text, nil
}

func convertStringPriceToFloatArray(text string) ([]float64, error) {

	stringPrices := strings.Split(text, " ")
	numberPrices := []float64{}
	var err error

	for _, val := range stringPrices {
		convertedPrice, parseError := strconv.ParseFloat(val, 64)

		if parseError != nil {
			err = errors.New("invalid character: please put number")
			numberPrices = []float64{}
			break
		}
		numberPrices = append(numberPrices, convertedPrice)

	}

	return numberPrices, err

}
func scanPrice() ([]float64, error) {

	fmt.Println("Please enter prices separated by space")

	text, err := readTextFromCML()

	if err != nil {
		return []float64{}, err
	}

	return convertStringPriceToFloatArray(text)

}

const priceFileName = "prices.txt"

func savePrices(prices []float64) {

	priceText := fmt.Sprint(prices)

	priceText = strings.Trim(priceText, "[]")

	fmt.Println(priceText)

	os.WriteFile(priceFileName, []byte(priceText), 0644)

}

func readPricesData() ([]float64, error) {
	data, err := os.ReadFile(priceFileName)

	if err != nil {
		return []float64{}, err
	}

	priceText := string(data)

	return convertStringPriceToFloatArray(priceText)
}
