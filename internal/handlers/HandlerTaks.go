// Package handlers
package handlers

import (
	"os"
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
	return CreateTaks(description)
}

func ListAll() (value []byte, err error) {
	return os.ReadFile("task/tasks.json")
}

func DeleteTask(ID int) string {
	return DeleteTasks(ID)
}
