package main

import "errors"

var InvalidInputError error = errors.New("Error: Please provide exactly one input file and one output file.")

func parseArgs(args []string) (string, string, error) {

	if len(args) > 3 || len(args) < 3 {
		return "", "", InvalidInputError
	}

	inputFileName := args[1]
	outputFileName := args[2]

	return inputFileName, outputFileName, nil
}
