package note_data

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Note struct {
	// field needs to be capitalized to be visible in note.json
	Title   string `json:"title"`
	Content string `json:"content"`
}

func New(title, content string) (*Note, error) {

	if title == "" || content == "" {
		return nil, errors.New("title and content are required")
	}
	return &Note{
		Title:   title,
		Content: content,
	}, nil

}

func (note *Note) SaveNote() error {

	fmt.Println(note)
	data, err := json.Marshal(note)

	fmt.Println(data)

	if err != nil {
		return err
	}

	err = os.WriteFile("note.json", data, 0644)

	if err != nil {
		return err
	}

	return nil

}
