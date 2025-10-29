// Package handlers
package handlers

import (
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

func CreateTask() {
	if VerificationDirectory("task") {
	} else {
	}
}

func ListAll() {
}

func DeleteTask(ID int) {
}
