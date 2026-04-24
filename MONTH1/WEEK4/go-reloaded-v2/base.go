package main

import (
	"strconv"
	"strings"
)

func ParseHex(content string) string {

	splitContent := strings.Fields(content)

	for index, word := range splitContent {
		if word == "(hex)" {

			hexVal := splitContent[index-1]

			newValue, err := strconv.ParseInt(hexVal, 16, 64)

			splitContent[index-1] = strconv.Itoa(int(newValue))

			splitContent = append(splitContent[:index], splitContent[index+1:]...)

		}
	}
	transformedHex := strings.Join(splitContent, " ")
	return transformedHex
}


func ParseBin(content string) string {

	splitContent := strings.Fields(content)

	for index, word := range splitContent {
		if word == "(bin)" {

			hexVal := splitContent[index-1]

			newValue, _ := strconv.ParseInt(hexVal, 2, 64)

			splitContent[index-1] = strconv.Itoa(int(newValue))

			splitContent = append(splitContent[:index], splitContent[index+1:]...)

		}
	}
	transformedBin := strings.Join(splitContent, " ")
	return transformedBin
}

func ParseOct(content string) string {

	splitContent := strings.Fields(content)

	for index, word := range splitContent {
		if word == "(oct)" {

			octVal := splitContent[index-1]

			newValue, _ := strconv.ParseInt(octVal, 8, 64)

			splitContent[index-1] = strconv.Itoa(int(newValue))

			splitContent = append(splitContent[:index], splitContent[index+1:]...)

		}
	}
	transformedOct := strings.Join(splitContent, " ")
	return transformedOct
}