package tax

var TaxRates = []int{0, 10, 20}

type newPriceWithTaxRate map[int][]float64

func (m newPriceWithTaxRate) Save() {

}

func New(currentPriceData []float64) *newPriceWithTaxRate {

	priceWithTax := make(newPriceWithTaxRate, 3)
	for _, tax := range TaxRates {
		for _, price := range currentPriceData {

			newPrice := price + (price * (float64(tax) / 100.00))

			priceWithTax[tax] = append(priceWithTax[tax], newPrice)

		}

	}

	return &priceWithTax
}
