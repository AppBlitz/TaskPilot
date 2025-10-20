package handlers

import "os"

// VarificationDirectory  verification if directory exist
func VarificationDirectory() bool {
	directories, _ := os.ReadDir(".")
	for _, files := range directories {
		if files.Name() == "task" {
			return true
		}
	}
	return false
}

// CreateDirectory implements creation directory if no exist
func CreateDirectory(name string) {
	pathAbsolute, _ := os.Getwd()
	os.Mkdir(pathAbsolute+"/"+name, 0o777)
}
