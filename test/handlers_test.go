package test

import (
	"testing"

	"github.com/AppBlitz/task_tracker/internal/handlers"
)

func TestUpdateTaskH(t *testing.T) {
	_, erro := handlers.DeleteTask(2)
	if erro != nil {
		t.Errorf("%s", "Erro deleted task")
	}
}

func TestUpdateTasksH(t *testing.T) {
	_, erro := handlers.UpdateTaks(3, "update task of handlers")
	if erro != nil {
		t.Errorf("%s", "Erro in update task")
	}
}
