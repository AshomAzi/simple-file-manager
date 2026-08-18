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
		content, err := ReadDirectory(path)
		dir := []os.DirEntry{}
		files := []os.DirEntry{}
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, value := range content {
			if value.IsDir() {
				dir = append(dir, value)
			} else {
				files = append(files, value)
			}
		}
		fmt.Printf("A total of %d content in the directory\n", len(content))
		fmt.Printf("Sub-directories: %d\n", len(dir))
		fmt.Printf("Files: %d\n", len(files))

		for _, val := range content {
			size, err := val.Info()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(val, size.Size())
		}

		for _, val := range dir {
			size, err := val.Info()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("📁 %s: %dbytes \n", val, size.Size())
		}
		for _, val := range files {
			size, err := val.Info()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("📄 %s: %dbytes \n",  val, size.Size())
		}
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

func ReadDirectory(path string) ([]os.DirEntry, error) {
	content, err := os.ReadDir(path)
	listOfContent := []os.DirEntry{}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, value := range content {
		listOfContent = append(listOfContent, value)
	}
	return listOfContent, nil
}
