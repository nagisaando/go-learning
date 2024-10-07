package main

import "fmt"

// NOTE: this is learning purpose. In real scenario, avoid using pointers for the small value like text or integers.
// Pointer can make code less readable
func main() {
	age := 32 // regular value

	var pointerOfAge *int     // asterisk mark for type declaration means it's a pointer type
	fmt.Println(pointerOfAge) // default value of pointer is <nil>
	pointerOfAge = &age

	editAgeAfter20Years(pointerOfAge)

	// this will prints the location of the memory
	fmt.Println(pointerOfAge)
	// if value that is stored in the address needs to be accessed, access with "*"
	fmt.Println(*pointerOfAge)

	fmt.Println(age)

}

func editAgeAfter20Years(age *int) {

	// this asterisk mark is dereferencing operator. It can access the value that is stored in the address
	*age += 20
}
