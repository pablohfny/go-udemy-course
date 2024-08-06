package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		title,
		content,
		time.Now(),
	}, nil
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	fileContent, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, fileContent, 0644)
}

func (note Note) Display() {
	fmt.Printf("Title: %v - %v\n", note.Title, note.CreatedAt.Format("2006-01-02"))
	fmt.Println(note.Content)
}
