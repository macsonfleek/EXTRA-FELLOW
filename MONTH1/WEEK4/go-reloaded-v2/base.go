package main

import (
	"strconv"
	"strings"
)

func ParseHex(content string) string {
	splitContent := strings.Fields(content)

	for index, word := range splitContent {
		if word == "(hex)" && index > 0 {
			hexVal := splitContent[index-1]
			newValue, err := strconv.ParseInt(hexVal, 16, 64)
			if err != nil {
				splitContent = append(splitContent[:index], splitContent[index+1:]...)
				continue
			}
			splitContent[index-1] = strconv.Itoa(int(newValue))
			splitContent = append(splitContent[:index], splitContent[index+1:]...)
		}
	}

	tranformedHex := strings.Join(splitContent, " ")
	return tranformedHex
}

func ParseBin(content string) string {
	splitContent := strings.Fields(content)

	for index, word := range splitContent {
		if word == "(bin)" && index > 0 {
			binVal := splitContent[index-1]
			newValue, err := strconv.ParseInt(binVal, 2, 64)
			if err != nil {
				splitContent = append(splitContent[:index], splitContent[index+1:]...)
				continue
			}
			splitContent[index-1] = strconv.Itoa(int(newValue))
			splitContent = append(splitContent[:index], splitContent[index+1:]...)
		}
	}

	tranformedBin := strings.Join(splitContent, " ")
	return tranformedBin
}

func ParseOct(content string) string {
	splitContent := strings.Fields(content)

	for index, word := range splitContent {
		if word == "(oct)" && index > 0 {
			octVal := splitContent[index-1]
			newValue, err := strconv.ParseInt(octVal, 8, 64)
			if err != nil {
				splitContent = append(splitContent[:index], splitContent[index+1:]...)
				continue
			}
			splitContent[index-1] = strconv.Itoa(int(newValue))
			splitContent = append(splitContent[:index], splitContent[index+1:]...)
		}
	}

	tranformedOct := strings.Join(splitContent, " ")
	return tranformedOct
}
