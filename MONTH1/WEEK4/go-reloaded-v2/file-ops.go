package main

import (
	"os"
)

func getFileContent(fileName string) (string, error) {

	byteData, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	fileContent := string(byteData)
	return fileContent, nil

}
