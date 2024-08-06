package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {

	if text == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		text,
	}, nil
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	fileContent, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, fileContent, 0644)
}

func (todo Todo) Display() {
	fmt.Printf(todo.Text)
}
