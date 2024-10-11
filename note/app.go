package main

import (
	"fmt"

	note_data "example.com/note/note-data"
)

func main() {

	var title string
	var content string

	createNoteData("enter title", &title)
	createNoteData("enter content", &content)

	note, err := note_data.New(title, content)

	if err != nil {
		panic(err)
	}

	err = note.SaveNote()

	if err != nil {
		panic(err)
	}

	fmt.Println("Saved!")

}

func createNoteData(outputText string, data *string) {
	fmt.Println(outputText)
	fmt.Scanln(data)
}
