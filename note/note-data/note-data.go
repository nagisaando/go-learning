package note_data

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Note struct {
	// field needs to be capitalized to be visible in note.json
	// otherwise json.Marshal can not catch it
	Title   string `json:"title"` // this is called struct tags (meta data) that can be used for field name for json file
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

func (note Note) Display() {
	fmt.Printf("You created note: \ntitle: %v\ncontent: %v\n", note.Title, note.Content)
}

func (note *Note) Save() error {

	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)

	// fmt.Println(note)
	data, err := json.Marshal(note) // convert data to json format and return byte[]

	// fmt.Println(data)

	if err != nil {
		return err
	}

	err = os.WriteFile(fileName+".json", data, 0644)

	if err != nil {
		return err
	}

	return nil

}
