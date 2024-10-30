package filemanager

import (
	"bufio"
	"encoding/json"
	"os"
)

type FileManager struct {
	InputPath  string
	OutputPath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputPath)

	if err != nil {
		return nil, err
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
		file.Close()
		return nil, err
	}

	file.Close() // we have to close the file when it is opened
	return lines, nil

}

func (fm FileManager) WriteResult(data interface{} /* or "any"*/) error {

	file, err := os.Create(fm.OutputPath)

	if err != nil {
		return err
	}

	// difference between json.marshal and json.encode [reference: https://stackoverflow.com/questions/33061117/in-golang-what-is-the-difference-between-json-encoding-and-marshalling]
	// [Key point]: they handle JSON encoding and output the result different

	// 1. json.Marshal:
	// - returns JSON byte slice
	// - we want to use it if we are plan to manipulate the data further, or send it over a network connection

	// 2. json.NewEncoder().Encode:
	// - writes directly to io.Writer (like os.Stdout, a file, or an HTTP response). It is convenient for streaming data or writing JSON output directly to a destination
	// - we want to use it if we want to write JSON output directly to an io.Writer, such as a file or HTTP response.

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return err
	}

	file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}
