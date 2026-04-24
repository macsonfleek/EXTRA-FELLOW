package main

import (
	"errors"
	"os"
)

var inputFileError error = errors.New("Error: Cannot read inpute file")

func getFileContent(filename string) (string, error) {
	byteData, err := os.ReadFile(filename)

	if err != nil {
		return "", inputFileError
	}
	fileContent := string(byteData)
	return fileContent, nil
}
