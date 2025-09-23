package handlers

import (
	"time"

	"github.com/AppBlitz/task_tracker/internal/model"
)

func NewTasks(id int, description string) *model.Tasks {
	return &model.Tasks{
		ID:          12,
		Description: description,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
}
