// Package handlers
package handlers

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/AppBlitz/task_tracker/internal/model"
)

const (
	COMPLETETASK model.StatusTaks = "done"
	PROGRESSTASK model.StatusTaks = "in-progress"
	STATUTODO    model.StatusTaks = "todo"
)

func NewTasks(id int, description string) *model.Tasks {
	return &model.Tasks{
		ID:          id,
		Description: description,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
		Status:      STATUTODO,
	}
}

func CreateTask(description string) string {
	idTaskCreate, erro := CreateTaks(description)
	if erro != nil {
		return ""
	}
	message := "Task added successfully" + "(ID:" + strconv.Itoa(idTaskCreate) + ")"
	return message
}

func ListAll(args []string) (value []byte, err error) {
	if len(args) > 0 {
		data := SearchForState(args[0])
		return json.Marshal(data)
	}
	return os.ReadFile("task/tasks.json")
}

func DeleteTask(ID int) (message string, erro error) {
	message, erro = DeleteTasks(ID)
	if erro != nil {
		return "", erro
	}
	return message, nil
}

func UpdateTaks(ID int, description string) (message string, erro error) {
	message = "Task update successfully with" + "(ID:" + strconv.Itoa(ID) + ")"
	erro = UpdateTask(ID, description)
	if erro != nil {
		return "", erro
	}
	return message, nil
}

func MarkDone(ID int) error {
	return MarkDones(ID)
}

func MarkProgress(ID int) {
	MarkInProgress(ID)
}
