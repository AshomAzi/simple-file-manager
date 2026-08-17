package main

import (
	"fmt"
	"os"
)

func main() {

	path, err := Path()
	if err != nil {
		fmt.Println("Invalid directory path!")
		return
	}
	fmt.Println(path)
}

func Path() (string, error) {

	inputPath := os.Args

	if len(inputPath) > 2 {
		return "Input only directory path", nil
	}
	var path string
	if len(inputPath) == 1 {
		defaultInput, err := os.Getwd()
		if err != nil {
			return "", err
		}
		path = defaultInput
	}

	if len(inputPath) == 2 {
		path = inputPath[1]
	}


	return path, nil
}

