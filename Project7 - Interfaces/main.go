package main

import (
	"bufio"
	"fmt"
	"notes/note"
	"notes/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type outputabble interface {
	saver
	Display()
}

// type outputabble interface {
// 	Save()
// 	Display()
// }

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	newNote, err := note.New(title, content)

	if err != nil {
		return
	}

	err = outputData(newNote)

	if err != nil {
		return
	}

	outputData(todo)
}

func getTodoData() string {
	return getUserInput("Todo text:")
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

func saveData(saver saver) error {
	err := saver.Save()

	if err != nil {
		fmt.Println("saving the todo failed.")
		return err
	}

	return nil
}

func outputData(data outputabble) error {
	data.Display()
	return saveData(data)
}

func useAnyData(data any) {
	switch data.(type) {
	case int:
		fmt.Println("Integer: ", data)
	case string:
		fmt.Println(data)
	}
}

func isInt(data any) (int, bool) {
	typedVal, ok := data.(int)
	return typedVal, ok
}

func generic[T int | float64 | string](a, b T) T {
	return a + b
}
