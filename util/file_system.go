package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FileSystem struct {
}

func (FileSystem) WriteToFile(filePath string, text string) {
	var f *os.File
	var err error

	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		f, err = os.Create(filePath)
	} else {
		f, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
	}

	gracefulExitIfError(err)

	defer f.Close()

	_, err = f.WriteString(text)
	gracefulExitIfError(err)
}

func (FileSystem) ReadFromFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func gracefulExitIfError(err error) {
	if err != nil {
		fmt.Printf("Inflight: %s\n", err.Error())
		os.Exit(0)
	}
}
