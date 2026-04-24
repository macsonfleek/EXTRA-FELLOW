package main

import "strings"

func ParseUp(content string) string {
	splitContent := strings.Fields(content)

	for index, word := range splitContent {
		if word == "(up)" && index > 0 {
			previousWord := splitContent[index-1]
			uppercaseWord := strings.ToUpper(previousWord)
			splitContent[index-1] = uppercaseWord
			splitContent = append(splitContent[:index], splitContent[index+1:]...)
		} else if word == "(up," && index > 0 {

		}
	}
	tranformedUp := strings.Join(splitContent, " ")
	return tranformedUp
}
