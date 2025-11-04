package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/AppBlitz/task_tracker/internal/handlers"
	"github.com/AppBlitz/task_tracker/internal/model"
	"github.com/charmbracelet/lipgloss/list"
)

func CreateResponseListTasks(args []string) {
	data, error := handlers.ListAll(args)
	if error != nil {
		log.Fatal(error)
	}
	value, _ := separateForState(data)
	fmt.Println(value)
}

func separateForState(value []byte) (l *list.List, erro error) {
	var auxiliary []model.Tasks
	var todo []byte
	var progress []byte
	var done []byte
	erro = json.Unmarshal(value, &auxiliary)
	if erro != nil {
		return nil, erro
	}
	for _, data := range auxiliary {
		switch data.Status {
		case handlers.COMPLETETASK:
			done = append(done, converByte(data)...)
		case handlers.PROGRESSTASK:
			progress = append(progress, converByte(data)...)
		default:
			todo = append(todo, converByte(data)...)
		}
	}
	response := list.New(
		"todo", list.New(string(todo)),
		"done", list.New(string(done)),
		"in-progress", list.New(string(progress)),
	)

	return response, nil
}

func converByte(tasks model.Tasks) []byte {
	value, _ := json.Marshal(tasks)
	return value
}

func VerificationNumber(number string) int {
	value, erro := strconv.Atoi(number)
	if erro != nil {
		return -1
	}
	return value
}
