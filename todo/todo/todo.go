package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string `json:"content"`
}

func New(content string) (*Todo, error) {

	if content == "" {
		return nil, errors.New("title and content are required")
	}
	return &Todo{
		Content: content,
	}, nil

}

func (note *Todo) Save() error {

	fmt.Println(note)
	data, err := json.Marshal(note) // convert data to json format and return byte[]

	fmt.Println(data)

	if err != nil {
		return err
	}

	err = os.WriteFile("todo.json", data, 0644)

	if err != nil {
		return err
	}

	return nil

}
