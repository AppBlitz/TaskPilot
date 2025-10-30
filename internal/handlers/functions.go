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

func returnId() {
}

func ReturnDataFile() (value []*model.Tasks, erro error) {
	data, erro := os.ReadFile(pathOfDirectory("task") + "/tasks.json")
	if erro != nil {
		return value, erro
	}
	erro = json.Unmarshal(data, &value)
	if erro != nil {
		return value, erro
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
