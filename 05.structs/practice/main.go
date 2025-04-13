package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"rishisingh.in/structs/note/note"
)

func main() {
	noteTitle, noteContent := getNoteData()

	userNote, err := note.New(noteTitle, noteContent)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Print()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}
	fmt.Println("Saving the note succeded")
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
