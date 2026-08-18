package actions

import (
	"fmt"
	"os"
)

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
