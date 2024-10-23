package price

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/price-calculator/scan"
)

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
func Scan() ([]float64, error) {

	fmt.Println("Please enter prices separated by space")

	text, err := scan.ReadTextFromCML()

	if err != nil {
		return []float64{}, err
	}

	return convertStringPriceToFloatArray(text)

}

const priceFileName = "prices.txt"

func Save(prices []float64) {

	priceText := fmt.Sprint(prices)

	priceText = strings.Trim(priceText, "[]")

	fmt.Println(priceText)

	os.WriteFile(priceFileName, []byte(priceText), 0644)

}

func Read() ([]float64, error) {
	data, err := os.ReadFile(priceFileName)

	if err != nil {
		return []float64{}, err
	}

	priceText := string(data)

	return convertStringPriceToFloatArray(priceText)
}
