package cmd

import (
	"fmt"
	"log"

	"github.com/AppBlitz/task_tracker/internal/handlers"
)

func CreateResponseListTasks() {
	data, error := handlers.ListAll()
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(string(data))
}
