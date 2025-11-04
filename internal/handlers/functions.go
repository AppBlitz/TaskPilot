package handlers

import (
	"encoding/json"
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

func CreateTaks(description string) (response int, erro error) {
	aux, _ := ReturnDataFile()
	var auxiliaryTasks []*model.Tasks
	if len(aux) != 0 {
		erro = json.Unmarshal(aux, &auxiliaryTasks)
		if erro != nil {
			panic(erro)
		}
	}
	_, erro = CreateFile()
	if erro != nil {
		return -1, erro
	}
	task := NewTasks(ReturnID(aux), description)
	erro = addData(orderTask(append(auxiliaryTasks, task)))
	if erro != nil {
		return -1, nil
	}
	return task.ID, nil
}

func ReturnDataFile() (value []byte, erro error) {
	var auxiliary []byte
	response, erro := pathOfDirectory("task")
	if erro != nil {
		return auxiliary, erro
	}
	return os.ReadFile(response + "/tasks.json")
}

func pathOfDirectory(nameDirectory string) (value string, erro error) {
	pathDirectory, erro := os.Getwd()
	if erro != nil {
		return "", erro
	}
	return pathDirectory + "/" + nameDirectory, nil
}

func CreateFile() (value bool, err error) {
	if !VerificationDirectory("task") {
		err = CreateDirectory("task")
		if err != nil {
			return false, err
		}
	}
	response, err := pathOfDirectory("task")
	if err != nil {
		return false, nil
	}
	_, err = os.Create(response + "/tasks.json")
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateDirectory implements creation directory if no exist
func CreateDirectory(nameDirectory string) error {
	pathAbsolute, _ := os.Getwd()
	erro := os.Mkdir(pathAbsolute+"/"+nameDirectory, 0o777)
	if erro != nil {
		return erro
	}
	return nil
}

func DeleteTasks(ID int) (message string, erro error) {
	aux, erro := ReturnDataFile()
	if erro != nil {
		return "", erro
	}
	var data []*model.Tasks
	var auxiliaryData []*model.Tasks
	if len(aux) != 0 {
		erro = json.Unmarshal(aux, &data)
		if erro != nil {
			return "", erro
		}
	}
	for _, value := range data {
		if value.ID != ID {
			auxiliaryData = append(auxiliaryData, value)
		}
	}
	erro = deleteFile()
	if erro != nil {
		return "", erro
	}
	erro = addData(auxiliaryData)
	if erro != nil {
		return "", erro
	}
	message = "task delete succesfully" + strconv.Itoa(ID)
	return message, nil
}

func addData(data []*model.Tasks) (erro error) {
	response, erro := pathOfDirectory("task")
	if erro != nil {
		return erro
	}
	file, erro := os.OpenFile(response+"/tasks.json", os.O_APPEND|os.O_RDWR, 0o0644)
	if erro != nil {
		return erro
	}
	valueByte, erro := json.Marshal(data)
	if erro != nil {
		return erro
	}
	_, erro = file.Write(valueByte)
	if erro != nil {
		return erro
	}
	return nil
}

func UpdateTask(ID int, description string) (erro error) {
	var datas []*model.Tasks
	var da []*model.Tasks
	auxData, _ := ReturnDataFile()
	erro = json.Unmarshal(auxData, &datas)
	if erro != nil {
		return erro
	}
	task := SearchTaksForID(datas, ID)
	for _, value := range datas {
		if value.ID != task.ID {
			da = append(da, value)
		}
		task.Description = description
		task.UpdateAt = time.Now()
	}
	erro = deleteFile()
	if erro != nil {
		return erro
	}
	erro = addData(orderTask(append(da, task)))
	if erro != nil {
		return erro
	}
	return nil
}

func deleteFile() (erro error) {
	erro = os.Remove("task/tasks.json")
	if erro != nil {
		return erro
	}
	_, erro = CreateFile()
	if erro != nil {
		return erro
	}
	return nil
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
		task.Status = COMPLETETASK
		erro = deleteFile()
		if erro != nil {
			panic(erro)
		}
		erro = addData(orderTask(append(datas, task)))
		if erro != nil {
			panic(erro)
		}
	}
}

func MarkInProgress(ID int) {
	var auxiliaryTask []*model.Tasks
	var datas []*model.Tasks
	data, erro := ReturnDataFile()
	if erro != nil {
		panic("erro read file")
	}
	erro = json.Unmarshal(data, &datas)
	if erro != nil {
		panic(erro)
	}

	task := SearchTaksForID(datas, ID)
	for _, value := range datas {
		if value.ID != task.ID {
			auxiliaryTask = append(auxiliaryTask, value)
		}
	}
	task.UpdateAt = time.Now()
	task.Status = PROGRESSTASK
	erro = deleteFile()
	if erro != nil {
		panic(erro)
	}
	erro = addData(orderTask(append(auxiliaryTask, task)))
	if erro != nil {
		panic(erro)
	}
}

func orderTask(tasks []*model.Tasks) []*model.Tasks {
	for i := range tasks {
		for j := i + 1; j < len(tasks); j += 1 {
			if tasks[i].ID > tasks[j].ID {
				auxiliary := tasks[j]
				tasks[j] = tasks[i]
				tasks[i] = auxiliary
			}
		}
	}
	return tasks
}
