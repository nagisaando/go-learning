package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	todo "example.com/todo/todo"
)

func main() {
	content := createTodo("enter content")

	note, err := todo.New(content)

	if err != nil {
		panic(err)
	}

	err = note.Save()

	if err != nil {
		panic(err)
	}

	fmt.Println("Saved!")

}

func createTodo(outputText string) string {
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
