package list

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
	// we can override the original array by re-assigning. and GO automatically clean up the memory space for the old array
	// also we can append multiple element at the same time
	prices = append(prices, 12345, 22222, 43333)

	discountPrices := []float64{12, 13, 14, 15}

	prices = append(prices, discountPrices...) // three dots unpack the array and merge the all elements
	fmt.Println(prices)

	// performance strategy
	// make(): this will reserve the memory space of 5 so when we add new stuff, GO does not go to allocate new space to create new array
	// 5: the number of memory to be reserved
	// 2: make array of length 2 and store null

	userName := make([]string, 2, 5)

	fmt.Println(userName) // it will show array with 2 null values

	userName[0] = "Nagisa" // already make() created userName array with 2 length so we can assign the value using index

	userName = append(userName, "Daniel") // this will be appended to index[2] (third element)

	fmt.Println(userName)
}

// ================================================================================

// // exercise:

// type Product struct {
// 	title string
// 	id    string
// 	price float64
// }

// func newProduct(title, id string, price float64) Product {
// 	return Product{
// 		title: title,
// 		id:    id,
// 		price: price,
// 	}
// }

// func main() {

// 	// 1) Create a new array (!) that contains three hobbies you have
// 	// 		Output (print) that array in the command line.
// 	// 2) Also output more data about that array:
// 	//		- The first element (standalone)
// 	//		- The second and third element combined as a new list
// 	// 3) Create a slice based on the first element that contains
// 	//		the first and second elements.
// 	//		Create that slice in two different ways (i.e. create two slices in the end)
// 	// 4) Re-slice the slice from (3) and change it to contain the second
// 	//		and last element of the original array.
// 	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 	// 7) Bonus: Create a "Product" struct with title, id, price and create a
// 	//		dynamic list of products (at least 2 products).
// 	//		Then add a third product to the existing list of products.

// 	// 1)
// 	var hobbies = [3]string{"Game", "Tennis", "Traveling"}

// 	// 2)
// 	// fmt.Println(hobbies[0])
// 	// fmt.Println(hobbies[1:])

// 	// 3)
// 	newHobbies := hobbies[0:2] // hobbies[:2]

// 	// 4)
// 	fmt.Println(newHobbies)
// 	fmt.Println(cap(newHobbies))
// 	fmt.Println(newHobbies[1:3])

// 	// 5)
// 	goals := []string{
// 		"being able to write a backend code",
// 		"being proficient in GO",
// 	}

// 	fmt.Println(goals)

// 	// 6)
// 	goals[1] = "New goal!"

// 	goals = append(goals, "Additional goal!")

// 	fmt.Println(goals)

// 	// 7)
// 	products := []Product{
// 		newProduct("ice tea", "123", 4.56),
// 		newProduct("coffee", "1232", 3),
// 	}

// 	fmt.Println(products)

// 	// or

// 	// GO knows the type without specifying Product inside the array
// 	products = []Product{
// 		{"ice tea", "123", 4.56},
// 		{"coffee", "1232", 3},
// 	}

// 	products = append(products, newProduct("milk", "100", 9))

// 	fmt.Println(products)
// }

// ================================================================================

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
