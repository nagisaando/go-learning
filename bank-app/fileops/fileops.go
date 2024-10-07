// if I want to create a new package, needs to be created within the subfolder. one go file is allowed per folder

package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// any variable, function that starts with uppercase is available to other packages (similar to "export" keyword in JS)
func WriteFloatFromFile(fileName string, value float64) {
	// Sprint() converts any type to human readable string.
	// For example 65 will be converted to "65"
	valueText := fmt.Sprint(value)

	// []byte(valueText) converts string to byte collection
	// 0644 means:
	// The file's owner can read and write (6)
	// Users in the same group as the file's owner can read (first 4)
	// All users can read (second 4)
	os.WriteFile(fileName, []byte(valueText), 0644)

}

func ReadFloatFromFile(fileName string, defaultReturnValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		// %2.f formats the value to 2 decimal places
		return defaultReturnValue, fmt.Errorf("failed to read a file, returning %2.f instead", defaultReturnValue)
	}

	// string() converts byte slices or unicode code points into their character/string representation
	// For example 65 will be converted to "A"
	valueText := string(data)

	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return defaultReturnValue, errors.New("failed to parse stored balance value")
	}

	return value, nil
}
