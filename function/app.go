package main

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double) // functions are first class values feature: it can be treated as a normal value as other types
	tripled := transformNumbers(&numbers, triple)

	// anonymous function
	quadruple := transformNumbers(&numbers, func(num int) int {
		return num * 4
	})

	quintuple := createTransformer(5) // [closure] 5 is now memorized and it will be used when the returned function is called

	fmt.Println(doubled)
	fmt.Println(tripled)
	fmt.Println(quadruple)
	fmt.Println(transformNumbers(&numbers, quintuple))

	transformFn := getTransformerFunction(&numbers)
	fmt.Println(transformNumbers(&numbers, transformFn))

	fmt.Println(factorial(3))
	fmt.Println(factorial(5))

	fmt.Println(sumUp(1, numbers...))
	fmt.Println(sumUp(1, 2, 3, 4))

}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
	transformedArr := []int{}

	for _, val := range *numbers {
		transformedArr = append(transformedArr, transform(val))
	}

	return transformedArr
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}

func getTransformerFunction(numbers *[]int) transformFunc {
	if (*numbers)[0] == 1 { // array pointer needs to be dereferenced explicitly unlike struct type
		return double // function can return another function as value
	} else {
		return triple
	}
}

// Factory function - function that returns other function with different config
func createTransformer(factor int) func(int) int {
	// Every anonymous function is a closure:
	// if the value from the outer scope is used in the function, the value gets saved and when the function is executed, the value is still saved and used.
	return func(num int) int {

		// function is a closure:
		// if the value from the outer scope is used in the function [factor], the value gets saved and
		// when the inner function is executed, the value from outer scope is still saved and used.
		return num * factor
	}
}

// recursive function:
// if num is 3:
// 	first loop returns 3 * factorial(2) and pauses the execution until factorial(2) is solved
// 	second loop returns 2 * factorial(1) and pauses the execution until factorial(1) is solved
//	third loop returns 1
// 	and now goes back to the second loop and calculates 2 * 1
// 	and goes back to the first loop and calculates 3 * 2
//	returns 6

func factorial(num int) int {
	if num == 1 {
		return 1
	}

	return num * factorial(num-1) // nested function needs to be executed and completed before the main function executes
}

// variadic functions: function that works with any number of parameters
// GO merge all the list of the parameter into a slice

func sumUp(startingValue int, numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}
