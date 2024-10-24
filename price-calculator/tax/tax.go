package tax

import (
	"encoding/json"
	"os"
)

var TaxRates = []int{0, 10, 20}

type newPriceWithTaxRate map[int][]float64

var fileName = "price_with_tax.json"

func (prices newPriceWithTaxRate) Save() error {

	data, err := json.Marshal(prices)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}
	return nil

}

func Read() (newPriceWithTaxRate, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var pricesData newPriceWithTaxRate

	// unmarshal can convert the type automatically
	// as long as the type specified matches json structure
	err = json.Unmarshal(data, &pricesData)

	if err != nil {
		return nil, err
	}

	return pricesData, nil

}

func New(currentPriceData []float64) newPriceWithTaxRate {

	priceWithTax := make(newPriceWithTaxRate, 3)
	for _, tax := range TaxRates {
		for _, price := range currentPriceData {

			newPrice := price + (price * (float64(tax) / 100.00))

			priceWithTax[tax] = append(priceWithTax[tax], newPrice)

		}

	}

	// maps are reference type and it returns reference without specifying pointer
	// struct is value type
	return priceWithTax
}
