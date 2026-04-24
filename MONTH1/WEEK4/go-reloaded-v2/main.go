package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// retrieve arguments from command line

	hexVal := "1FFF"

	NewHexValue, _ := strconv.ParseInt(hexVal, 16, 64)

	fmt.Println(NewHexValue)

	args := os.Args

	inputFileName, _, err := parseArgs(args)

	if err != nil {
		fmt.Println(err)
		return
	}

	inputeContent, err := getFileContent(inputFileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Hex Transformation
	transformedHex := ParseHex(inputeContent)
	fmt.Println(transformedHex)

	// Bin Transformation
	transformedBin := ParseBin(transformedHex)
	fmt.Println(transformedBin)

	// Oct Transformation
	transformedOct := ParseOct(transformedBin)
	fmt.Println(transformedOct)


}

