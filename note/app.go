package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	note_data "example.com/note/note-data"
	todo "example.com/note/todo"
)

// interface = contract that certain value or struct has a certain method
// name convention ["the name of method" + "er"]: when creating interface that contains only one method
type saver interface {
	// interface does not define the method logic but just to tell which method is available
	Save() error
}

// embedded interfaces
type outputtable interface {
	saver
	Display()
}

func main() {

	title := createData("enter title")
	content := createData("enter content")

	todoText := createData("enter todo content")

	todo, err := todo.New(todoText)

	if err != nil {
		panic(err)
	}

	note, err := note_data.New(title, content)

	if err != nil {
		panic(err)
	}

	err = outputData(todo, "Todo Saved!")

	if err != nil {
		panic(err)
	}

	err = outputData(note, "Note Saved!")
	if err != nil {
		panic(err)
	}

}

func printSomething(value interface{}) { // interface{} or "any" indicates any value type is accepted
	fmt.Println(value)
}

func outputData(data outputtable, successMessage string) error {
	data.Display()
	return saveData(data, successMessage) // returning saveData will return error automatically
}
func saveData(data saver, successMessage string) error { // [data save]r: accept the data that signed the saver interface contract: the value that has save()
	err := data.Save()

	if err != nil {
		return err
	}

	fmt.Println(successMessage)

	return nil
}
func createData(outputText string) string {
	fmt.Println(outputText)
	// fmt.Scanln(data) // Scanln() works only single word or number (does not work with space)

	// bufio works to read long text
	reader := bufio.NewReader(os.Stdin) // making struct and targeting command line to scan text

	// '\n' => this means when user hit "enter" key, it stop scanning text
	// '\n' => to specify single byte, single character, needs to be wrapped with single quotes
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // for windows since return key return as \r\n

	return text
}
