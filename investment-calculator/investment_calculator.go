package main

import (
	"fmt"
	"math"
)

func main() {
  var investmentAmount, expectedReturnRate, years float64
  var futureValue float64

  fmt.Println("Investment amount: ") // after prints the text, move to the new line
  fmt.Scan(&investmentAmount)

  fmt.Print("Expected return rate: ")
  fmt.Scanln(&expectedReturnRate)

  fmt.Print("Years: ")
  fmt.Scanln(&years)

 
  totalReturnRate := 1.0 + (float64(expectedReturnRate) / 100.0)

  fmt.Print(totalReturnRate)
  // need to check what is Pow
  futureValue = float64(investmentAmount) * math.Pow(totalReturnRate, float64(years))
  fmt.Print("Future value:", investmentAmount, "and", expectedReturnRate, years, futureValue)
}