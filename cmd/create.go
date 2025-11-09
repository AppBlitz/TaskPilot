package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/AppBlitz/task_tracker/internal/handlers"
)

func CreateResponseListTasks(args []string) {
	data, error := handlers.ListAll(args)
	if error != nil {
		log.Fatal(error)
	}
	style := []string{"#6A5ACD", "#CC99FF", "#1A1A1A"}
	fmt.Println(handlers.Styles("list all tasks", len(data), style).Render(string(data)))
}

func VerificationNumber(number string) int {
	value, erro := strconv.Atoi(number)
	if erro != nil {
		return -1
	}
	return value
}

func CommandCreateTasks(description string) {
	style := []string{"#6A5ACD", "#cc99ff", "#1A1A1A"}
	message := handlers.CreateTask(description)
	fmt.Println(handlers.Styles("", len(message), style).Render(message))
}
