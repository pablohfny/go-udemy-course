package main

import (
	"bufio"
	"fmt"
	"notes/note"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()

	newNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	newNote.Display()

	err = newNote.Save()

	if err != nil {
		fmt.Println("saving the note failed.")
		return
	}

	fmt.Println("saving the note succeeded.")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}

	value = strings.TrimSuffix(value, "\r")
	value = strings.TrimSuffix(value, "\n")

	return value
}
