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
	validatedPath, err := ValidatePath(path)
	if err != nil {
		fmt.Println("Invalid Directory or directory does not exist!")
		return
	}
	if validatedPath == true {
		content, err := os.ReadDir(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(content)
	}
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

func ValidatePath(path string) (bool, error) {

	_, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return true, nil
}
