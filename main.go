package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/structs-practice/note"
)

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}

	noteStruct, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	noteStruct.WriteFile()
	noteContentFromFile, err := noteStruct.ReadFile(title)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(noteContentFromFile)
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Title: ")
	if err != nil {
		return "", "", err
	}
	content, err := getUserInput("Content: ")
	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	// var value string
	// fmt.Scanln(&value)

	// return value
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return "", errors.New("reader error")
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text, nil
}
