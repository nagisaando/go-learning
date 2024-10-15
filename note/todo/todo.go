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

func (todo Todo) Display() {
	fmt.Printf("You created todo list: %v\n", todo.Content)
}

func (todo *Todo) Save() error {

	// fmt.Println(todo)
	data, err := json.Marshal(todo) // convert data to json format and return byte[]

	// fmt.Println(data)

	if err != nil {
		return err
	}

	err = os.WriteFile("todo.json", data, 0644)

	if err != nil {
		return err
	}

	return nil

}
