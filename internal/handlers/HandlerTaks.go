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
		Status:      PROGRESSTASK,
	}
}

func CreateTask(description string) string {
	idTaskCreate := CreateTaks(description)
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

func DeleteTask(ID int) string {
	return DeleteTasks(ID)
}

func UpdateTaks(ID int, description string) string {
	message := "Task update successfully with" + "(ID:" + strconv.Itoa(ID) + ")"
	if !UpdateTask(ID, description) {
		message = "Task no update with succesfully"
	}
	return message
}
