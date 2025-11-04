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
	var auxiliary []*model.Tasks
	var todo []*model.Tasks
	var progress []*model.Tasks
	var done []*model.Tasks
	erro = json.Unmarshal(value, &auxiliary)
	if erro != nil {
		return nil, erro
	}
	for _, data := range auxiliary {
		switch data.Status {
		case handlers.COMPLETETASK:
			done = append(done, data)
		case handlers.PROGRESSTASK:
			progress = append(progress, data)
		default:
			todo = append(todo, data)
		}
	}
	l = list.New(
		"todo", list.New(todo),
		"done", list.New(done),
		"in-progress", list.New(progress),
	)
	return l, nil
}

func VerificationNumber(number string) int {
	value, erro := strconv.Atoi(number)
	if erro != nil {
		return -1
	}
	return value
}
