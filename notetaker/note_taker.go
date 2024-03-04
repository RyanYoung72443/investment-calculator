package notetaker

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"training/interfaces"
	"training/notetaker/note"
	"training/notetaker/todo"
)

func App(file string) {
	var err error

	if file == "note" {
		title, content := getNoteData()
		err = handleReturn(note.New(title, content))
	}

	if file == "todo" {
		todoText := getUserInput("Todo text:")
		err = handleReturn(todo.New(todoText))
	}

	if err != nil {
		fmt.Println(err)
		return
	}
}

func handleReturn(value interfaces.Outputable, err error) error {
	if err != nil {
		return err
	}
	return outputData(value)
}

func outputData(data interfaces.Outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data interfaces.Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
