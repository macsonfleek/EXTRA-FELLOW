package main

import (
	"fmt"
	"os"
)

func main() {
	// Retrieve all arguments from the command line
	args := os.Args

	// Getting Input and Output Filename
	inputFileName, _, err := parseArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the content of the Input Filename
	inputContent, err := getFileContent(inputFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// HEX TRANSFORMATION
	transformedHex := ParseHex(inputContent)

	// BIN TRANSFORMATION
	transformedBin := ParseBin(transformedHex)

	// OCT TRANSFORMATION
	tranformedOct := ParseOct(transformedBin)

	// UP TRANFORMATION
	tranformedUp := ParseUp(tranformedOct)
	fmt.Println(tranformedUp)

}
