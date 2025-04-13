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
		Text: text,
	}, nil
}

func (todo Todo) Print() {
	fmt.Println("Todo:", todo.Text)
}

func (todo Todo) Save() error {
	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile("todo.json", json, 0644)
}
