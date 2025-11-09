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
	style := []string{"#6A5ACD", "#1A1A1A", "#CC99FF"}
	fmt.Println(handlers.Styles("list all tasks", style).Render(string(data)))
}

func VerificationNumber(number string) int {
	value, erro := strconv.Atoi(number)
	if erro != nil {
		return -1
	}
	return value
}

func CommandCreateTasks(description string) {
	style := []string{"#ccff99", "#cc99ff", "#cccc66"}
	message := handlers.CreateTask(description)
	fmt.Println(handlers.Styles("", style).Render(message))
}
