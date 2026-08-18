package main

import (
	"fmt"
	p "manager/actions"
	"os"
)

func main() {

	path, err := p.Path()
	if err != nil {
		fmt.Println("Invalid directory path!")
		return
	}
	validatedPath, err := p.ValidatePath(path)
	if err != nil {
		fmt.Println("Invalid Directory or directory does not exist!")
		return
	}
	if validatedPath == true {
		content, err := p.ReadDirectory(path)
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

		for _, val := range dir {
			_, err := val.Info()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("📁 %s \n", val)
		}
		for _, val := range files {
			size, err := val.Info()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("📄 %s: %dbytes \n", val, size.Size())
		}
	}
}
