// Package test are test project
package test

import (
	"encoding/json"
	"testing"

	"github.com/AppBlitz/task_tracker/internal/handlers"
	"github.com/AppBlitz/task_tracker/internal/model"
)

// func TestCreateFile(t *testing.T) {
// 	_, erro := handlers.CreateFile()
// 	if erro != nil {
// 		t.Errorf("%s", "creation file no succesfull")
// 	}
// }

func TestReturnDataFile(t *testing.T) {
	_, erro := handlers.ReturnDataFile()
	if erro != nil {
		t.Errorf("%s", "File no found")
	}
}

func TestSearchTaskForID(t *testing.T) {
	value, erro := handlers.ReturnDataFile()
	var auxiliar []*model.Tasks
	if erro != nil {
		t.Errorf("%s", "file not found")
	}
	erro = json.Unmarshal(value, &auxiliar)
	if erro != nil {
		t.Errorf("%s %s", "file parser erro", "file save task is empty")
	}
}

func TestCreateTask(t *testing.T) {
	description := "task for test"
	task := handlers.CreateTask(description)
	if task == "" {
		t.Errorf("%s", "Handler no create with succesfull")
	}
}

func TestMarkDone(t *testing.T) {
	erro := handlers.MarkDone(2)
	if erro != nil {
		t.Errorf("%s", "error in update task to done")
	}
}

func TestMakProgress(t *testing.T) {
	erro := handlers.MarkInProgress(2)
	if erro != nil {
		t.Errorf("%s", "error in update task to progress")
	}
}

func TestUpdateTask(t *testing.T) {
	_, erro := handlers.UpdateTaks(1, "task modification test")
	if erro != nil {
		t.Errorf("%s", "Error in update task")
	}
}

func TestDeleteTask(t *testing.T) {
	_, erro := handlers.DeleteTasks(1)
	if erro != nil {
		t.Errorf("%s", "erro delete task")
	}
}
