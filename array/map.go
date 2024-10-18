package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}
func main() {

	// map: similar to array: group of elements that are stored as key-value pairs
	websites := map[string]string{ // [string] => type of key, string => type of value
		"Udemy":  "https://www.udemy.com/",
		"Github": "https://github.com/",
	}

	websites["Youtube"] = "https://www.youtube.com/"

	delete(websites, "Udemy")

	fmt.Println(websites)

	// make() in maps, unlike with slice, it does not pre-populate elements with null but reserve the memory space

	courseRating := make(floatMap, 3)

	courseRating["udemy"] = 4.5
	courseRating["react"] = 5

	courseRating.output()

	for key, value := range courseRating {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
	}

}
