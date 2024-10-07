package main

import (
	"fmt"
	"os"
)

func main() {
	var revenue float64
	var expense float64
	var taxRate float64

	// this is alternative way to pass the value itself instead of memory address
	// revenue = scanText("Enter revenue: ")
	// expense = scanText("Enter expense: ")
	// taxRate = scanText("Enter tax rate: ")

	scanText("Enter revenue: ", &revenue) // & is used to get the memory address of a variable: where the value is stored
	scanText("Enter expense: ", &expense)
	scanText("Enter tax rate: ", &taxRate)

	fmt.Printf("revenue: %v \n", revenue)
	fmt.Printf("expense: %v \n", expense)
	fmt.Printf("taxRate: %v \n", taxRate)
	earningsBeforeTax := revenue - expense
	taxPaid, profit := calculateTaxPaidAndProfit(earningsBeforeTax, taxRate)

	fmt.Printf("You have paid the tax of the amount: %.2f\n", taxPaid)
	writeResultToFile("tax-paid.txt", taxPaid)

	ratio := earningsBeforeTax / profit

	// Sprintf can format the value and return string, it can be stored as variable
	formattedEarningsBeforeTaxSentence := fmt.Sprintf("Your profit: %.2f\n", profit)

	// fmt.Printf formats string
	fmt.Printf("Your EBT: %v\n", earningsBeforeTax)
	writeResultToFile("ebt.txt", earningsBeforeTax)

	// shows only one decimal place
	// fmt.Printf("Your profit: %.1f\n", profit)

	// fmt.Println(formattedEarningsBeforeTaxSentence)
	outputText(formattedEarningsBeforeTaxSentence)
	fmt.Println("The ratio: ", ratio)
	writeResultToFile("ratio.txt", ratio)

}

func writeResultToFile(fileName string, amount float64) {
	amountString := fmt.Sprint(amount)
	err := os.WriteFile(fileName, []byte(amountString), 0644)
	if err != nil {
		panic(err)
	}
}
func outputText(text string) {
	fmt.Print(text)
}

func scanText(outputText string, scanningValue *float64) {
	fmt.Print(outputText)
	_, err := fmt.Scanln(scanningValue)
	if err != nil {
		panic("invalid input. please try again")

	}
	if *scanningValue < 0 { // * is dereferencing operator: it gives you an access to the value stored at the memory address the pointer is pointing to
		panic("please don't put negative number or non-number value")
	}
}

// this is alternative way to scanText()

// revenue = scanText("Enter revenue: ")
// expense = scanText("Enter expense: ")
// taxRate = scanText("Enter tax rate: ")

// func anotherScanText(outputText string) float64 {
// 	fmt.Print(outputText)
// 	var userInput float64
// 	_, err := fmt.Scanln(&userInput) // scanln() returns 1 if scanning succeed
// 	if err != nil {
// 		panic("invalid input. please try again")

// 	}
// 	if userInput < 0 { // * is dereferencing operator: it gives you an access to the value stored at the memory address the pointer is pointing to
// 		panic("please don't put negative number or non-number value")
// 	}

// 	return userInput
// }

func calculateTaxPaidAndProfit(earningsBeforeTax, taxRate float64) (taxPaid float64, profit float64) { // you can declare the value as return value

	taxPaid = earningsBeforeTax * (taxRate / 100)
	fmt.Println(earningsBeforeTax)
	fmt.Println(taxPaid)
	fmt.Println(taxRate)
	fmt.Println(taxRate / 100)
	profit = earningsBeforeTax - taxPaid

	return taxPaid, profit // go can return two values!!!

}
