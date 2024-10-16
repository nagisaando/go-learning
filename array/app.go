package main

import "fmt"

func main() {
	// without specifying the length of array, GO create a SLICE instead of array
	// and creates the array to the slice behind the scenes.
	// this is needed to create dynamic length of array
	prices := []float64{12, 23}

	// this will append element to the array (slice)
	// To add new element, Go will create new brand array based on the original one behind the scene (because the array length is fixed, can not change it dynamically)
	// and add the new element to the recently created array since the array has the new capacity.
	updatedPrices := append(prices, 12345) // this will not change the original array
	fmt.Println(updatedPrices, prices)
	prices = append(prices, 12345) // we can override the original array by re-assigning. and GO automatically clean up the memory space for the old array

	fmt.Println(prices)
}

// func main() {
// 	prices := [4]float64{1, 2.23, 3, 4.4}
// 	fmt.Println(prices)

// 	var product [3]string = [3]string{"Product 1"}

// 	fmt.Println(product) // it'll show empty for the array item that has not contain any element

// 	product[2] = "product 3"
// 	fmt.Println(product[2])

// 	// [Array slice]
// 	// this means from index 1 to index 3 BUT index 3 is not extracted.
// 	// so it will be from index 1 to 2
// 	featuredPrice := prices[1:3]

// 	// featuredPrice2 := prices[:3] // from the beginning of array to index 2 (since 3 is not included)
// 	// featuredPrice3 := prices[1:] // from index 1 to the end of the array

// 	fmt.Println(featuredPrice)
// 	// fmt.Println(featuredPrice2)
// 	// fmt.Println(featuredPrice3)

// 	// this will return 2, 3
// 	// length 2: because featuredPrice is [2.23 3] and has length of 2
// 	// cap 3: because it shows number of the available elements of the original array (prices)
// 	// 		  featuredPrice still can access 4.4 (last element of prices) if they want to

// 	fmt.Println(len(featuredPrice), cap(featuredPrice))

// 	featuredPrice = featuredPrice[:3] // still can access 4.4 because it has a capacity
// 	fmt.Println(featuredPrice)

// 	// changing element through sliced array also changes the original array
// 	// featuredPrice[0] = 10000
// 	// fmt.Println(prices)
// }
