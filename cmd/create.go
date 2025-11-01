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
	fmt.Println(string(data))
}

func VerificationNumber(number string) int {
	value, erro := strconv.Atoi(number)
	if erro != nil {
		return -1
	}
	return value
}
