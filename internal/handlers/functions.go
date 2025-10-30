package handlers

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/AppBlitz/task_tracker/internal/model"
)

// VerificationDirectory  verification if directory exist
func VerificationDirectory(directoryName string) bool {
	directories, _ := os.ReadDir(".")
	for _, files := range directories {
		if files.Name() == directoryName {
			return true
		}
	}
	return false
}

func ReturnID() int {
	var auxiliaryTasks []*model.Tasks
	var auxiliaryID int
	value, erro := os.ReadFile(pathOfDirectory("task") + "/tasks.json")
	if erro != nil {
		panic(erro)
	}
	if len(value) > 0 {
		erro = json.Unmarshal(value, &auxiliaryTasks)
		for index, data := range auxiliaryTasks {
			if index == len(auxiliaryTasks) {
				auxiliaryID = data.ID
			}
		}
		if erro != nil {
			panic(erro)
		}
	}
	return auxiliaryID + 1
}

func CreateTaks(description string) string {
	file, erro := os.OpenFile("task/tasks.json", os.O_RDWR|os.O_APPEND, 0o0644)
	if erro != nil {
		panic(erro)
	}
	defer file.Close()
	_, erro = CreateFile()
	if erro != nil {
		panic(erro)
	}
	var message string
	task := NewTasks(ReturnID(), description)
	data, erro := json.Marshal(task)
	if erro != nil {
		panic(erro)
	}
	_, erro = file.Write(data)
	if erro != nil {
		panic(erro)
	}

	return message
}

func ReturnDataFile() (value []*model.Tasks, erro error) {
	data, erro := os.ReadFile(pathOfDirectory("task") + "/tasks.json")
	if erro != nil {
		return value, erro
	}
	if len(value) > 0 {
		erro = json.Unmarshal(data, &value)
		if erro != nil {
			return value, erro
		}
	}
	return value, nil
}

func pathOfDirectory(nameDirectory string) (value string) {
	pathDirectory, erro := os.Getwd()
	if erro != nil {
		return ""
	}
	return pathDirectory + "/" + nameDirectory
}

func CreateFile() (value bool, err error) {
	if !VerificationDirectory("task") {
		CreateDirectory("task")
	}
	_, erro := os.Create(pathOfDirectory("task") + "/tasks.json")
	if erro != nil {
		return false, erro
	}
	return true, nil
}

// CreateDirectory implements creation directory if no exist
func CreateDirectory(nameDirectory string) {
	pathAbsolute, _ := os.Getwd()
	err := os.Mkdir(pathAbsolute+"/"+nameDirectory, 0o777)
	if err != nil {
		panic(err)
	}
}

func DeleteTasks(ID int) (message string) {
	data, erro := ReturnDataFile()
	var auxiliaryData []*model.Tasks
	if erro != nil {
		log.Fatal(erro)
	}
	for _, value := range data {
		if value.ID != ID {
			auxiliaryData = append(auxiliaryData, value)
		}
	}
	erro = os.Remove(pathOfDirectory("task") + "/tasks.json")
	verification, _ := CreateFile()
	if erro == nil && verification {
		if addData(auxiliaryData) {
			message = "task delete succesfully" + strconv.Itoa(ID)
		}
	}
	return message
}

func addData(data []*model.Tasks) bool {
	file, erro := os.OpenFile(pathOfDirectory("task")+"/tasks.json", os.O_APPEND|os.O_RDWR, 0o0644)
	if erro != nil {
		panic("Error read file")
	}
	defer file.Close()
	for _, value := range data {
		valueByte, erro := json.Marshal(value)
		if erro != nil {
			panic("erro run loop")
		}
		_, erro = file.Write(valueByte)
		if erro != nil {
			panic("error add value")
		}
	}
	return true
}
