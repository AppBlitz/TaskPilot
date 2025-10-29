package handlers

import (
	"os"
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
