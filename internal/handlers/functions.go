package handlers

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

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

func ReturnID(aux []byte) int {
	var auxiliaryTasks []*model.Tasks
	if len(aux) != 0 {
		erro := json.Unmarshal(aux, &auxiliaryTasks)
		if erro != nil {
			panic(erro)
		}
	}
	if len(auxiliaryTasks) == 0 {
		return 1
	}
	return auxiliaryTasks[len(auxiliaryTasks)-1].ID + 1
}

func CreateTaks(description string) int {
	aux, _ := ReturnDataFile()
	var auxiliaryTasks []*model.Tasks
	if len(aux) != 0 {
		erro := json.Unmarshal(aux, &auxiliaryTasks)
		if erro != nil {
			panic(erro)
		}
	}
	_, erro := CreateFile()
	if erro != nil {
		panic(erro)
	}
	task := NewTasks(ReturnID(aux), description)
	auxiliaryTasks = append(auxiliaryTasks, task)
	addData(auxiliaryTasks)
	return task.ID
}

func ReturnDataFile() (value []byte, erro error) {
	return os.ReadFile(pathOfDirectory("task") + "/tasks.json")
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
	aux, _ := ReturnDataFile()
	var data []*model.Tasks
	var auxiliaryData []*model.Tasks
	if len(aux) != 0 {
		erro := json.Unmarshal(aux, &data)
		if erro != nil {
			log.Fatal(erro)
		}
	}
	for _, value := range data {
		if value.ID != ID {
			auxiliaryData = append(auxiliaryData, value)
		}
	}
	erro := os.Remove(pathOfDirectory("task") + "/tasks.json")
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
	valueByte, erro := json.Marshal(data)
	if erro != nil {
		panic(erro)
	}
	_, erro = file.Write(valueByte)
	if erro != nil {
		panic(erro)
	}
	return true
}

func UpdateTask(ID int, description string) bool {
	var datas []*model.Tasks
	var da []*model.Tasks
	auxData, _ := ReturnDataFile()
	erro := json.Unmarshal(auxData, &datas)
	if erro != nil {
		return false
	}
	task := SearchTaksForID(datas, ID)
	for _, value := range datas {
		if value.ID != task.ID {
			da = append(da, value)
		}
		task.Description = description
		task.UpdateAt = time.Now()
	}
	deleteFile()
	addData(append(da, task))
	return true
}

func deleteFile() {
	erro := os.Remove("task/tasks.json")
	if erro != nil {
		panic("Erro delete file")
	}
	_, erro = CreateFile()
	if erro != nil {
		panic("Erro create file")
	}
}

func SearchTaksForID(auxiliaryTask []*model.Tasks, ID int) *model.Tasks {
	var taskFound *model.Tasks
	for _, value := range auxiliaryTask {
		if value.ID == ID {
			taskFound = value
		}
	}
	return taskFound
}

func SearchForState(state string) []*model.Tasks {
	var auxilairyData []*model.Tasks
	var auxilairyDatas []*model.Tasks
	data, erro := ReturnDataFile()
	if erro != nil {
		panic("erro read file")
	}
	erro = json.Unmarshal(data, &auxilairyDatas)
	if erro != nil {
		panic("Error converter value")
	}
	for _, value := range auxilairyDatas {
		if verificationState(model.StatusTaks(state), value) {
			auxilairyData = append(auxilairyData, value)
		}
	}
	return auxilairyData
}

func verificationState(state model.StatusTaks, taskVerification *model.Tasks) bool {
	return state == taskVerification.Status
}

func MarkDones(ID int) {
	var auxiliaryTask []*model.Tasks
	var datas []*model.Tasks
	data, erro := ReturnDataFile()
	if erro != nil {
		panic(erro)
	}
	erro = json.Unmarshal(data, &auxiliaryTask)
	if erro != nil {
		panic(erro)
	}
	task := SearchTaksForID(auxiliaryTask, ID)
	for _, value := range auxiliaryTask {
		if value.ID != task.ID {
			datas = append(datas, value)
		}
		task.UpdateAt = time.Now()
		task.Status = "done"
		deleteFile()
		addData(append(datas, task))
	}
}
