package main

import "fmt"

func main() {

	// map: similar to array: group of elements that are stored as key-value pairs
	websites := map[string]string{ // [string] => type of key, string => type of value
		"Udemy":  "https://www.udemy.com/",
		"Github": "https://github.com/",
	}

	websites["Youtube"] = "https://www.youtube.com/"

	delete(websites, "Udemy")

	fmt.Println(websites)

}
