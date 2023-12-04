package main

import (
	"fmt"
	"os"
	"path/filepath"

	"adventofcode.2023.goloang.chapman/day1"
)

func main() {
	c("1", true)
}

func c(day string, useFullInput bool) {
	fullInput, smallInput := getInput(day)
	var input string
	if useFullInput {
		input = fullInput
	} else {
		input = smallInput
	}

	switch day {
	case "1":
		day1.Execute(input)
	default:
		fmt.Println("Invalid day for input: " + input)
	}
}

func getInput(day string) (string, string) {
	path := getPath()

	inputPath := path + string(os.PathSeparator) + "input" + string(os.PathSeparator) + "day" + day + ".input"
	smallnput := inputPath + ".small"

	return inputPath, smallnput
}

func getPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}
