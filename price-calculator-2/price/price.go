package price

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{0, 10, 20},
		TaxRate:     taxRate,
		// we don't initialize taxIncludedPrices at the moment since it needs to be calculated later
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		// fmt.Sprintf can return float64 to string

		taxIncludedPrice := price * (1 + (job.TaxRate / 100))
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)

}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("price.txt")

	if err != nil {
		fmt.Print(err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	// scanner.Scan() read one line at a time.
	// and it returns false if there is no more scannable text
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // We can get the scanned text by scanner.Text()
	}

	// we can detect if there is any scan error by calling scanner.Err()
	err = scanner.Err()

	if err != nil {
		fmt.Println(err)
		file.Close() // we have to close the file when it is opened
		return
	}

	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println(err)
			file.Close()

			return

		}

		prices[lineIndex] = floatPrice
	}

	// this update won't work if we don't get job data as pointer
	// works: func (job *TaxIncludedPriceJob)
	// does NOT work: func (job TaxIncludedPriceJob) since it is just a copy and can not update the original data

	job.InputPrices = prices

}
