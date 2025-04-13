package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"rishisingh.in/structs/note/note"
	"rishisingh.in/structs/note/todo"
)

type saver interface {
	Save() error
}

// type printer interface {
// 	Print()
// }

type outputable interface {
	saver
	Print()
}

// type outputable interface {
// 	Save() error
// 	Print()
// }

func main() {
	noteTitle, noteContent := getNoteData()
	todoText := getUserInput("Enter Todo: ")

	todo, err := todo.New(todoText)
	if err != nil {
		printSomething(err)
		return
	}

	userNote, err := note.New(noteTitle, noteContent)
	if err != nil {
		printSomething(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	outputData(userNote)
	// if err != nil {
	// 	return
	// }
}

func printSomething(data interface{}) {
	intVal, ok := data.(int)
	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := data.(float64)
	if ok {
		fmt.Println("Float64:", floatVal)
		return
	}

	stirngVal, ok := data.(string)
	if ok {
		fmt.Println("String:", stirngVal)
		return
	}
	// switch data.(type) {
	// case string:
	// 	fmt.Println("String:", data)
	// case int:
	// 	fmt.Println("Integer:", data)
	// case float64:
	// 	fmt.Println("Float64:", data)
	// default:
	// 	fmt.Println(data)
	// }
}

func outputData(data outputable) error {
	data.Print()

	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		printSomething("Saving data failed")
		return err
	}
	printSomething("Saving data succeded")
	return nil
}

func getNoteData() (string, string) {
	noteTitle := getUserInput("Enter Note Title: ")
	noteContent := getUserInput("Enter Note Content: ")

	return noteTitle, noteContent
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
