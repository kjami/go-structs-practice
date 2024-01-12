package note

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"example.com/structs-practice/file"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		n := Note{}
		return &n, errors.New("title and content should not be empty")
	}
	n := Note{
		title,
		content,
		time.Now(),
	}
	return &n, nil
}

func (note Note) WriteFile() error {
	fileName := getFileName(note.Title)
	jsonContent, err := json.MarshalIndent(note, "", "   ")
	if err != nil {
		return err
	}
	return file.WriteFile(fileName, jsonContent)
}

func (Note) ReadFile(title string) (string, error) {
	fileName := getFileName(title)
	return file.ReadFile(fileName)
}

func getFileName(title string) string {
	fileName := strings.ToLower(title)
	fileName = strings.ReplaceAll(fileName, " ", "_")
	return fileName + ".json"
}
