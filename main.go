package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func writeFile(fileName string, content string) {
	fileContent := []byte(content)
	err := os.WriteFile(fileName, fileContent, 0644)
	check(err)
}

func copyFile(filePath string, newFilepath string) {
	result, err := os.ReadFile(filePath)
	check(err)
	writeFile(newFilepath, string(result))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Project name is required, run the cli and <project-name>")
		return
	}

	dirName := "./" + os.Args[1]
	fmt.Println("Starting project " + os.Args[1])
	os.Mkdir(dirName, 0644)

	copyFile("./templates/index.html", "./"+dirName+"/"+"index.html")
	copyFile("./templates/index.css", "./"+dirName+"/"+"index.css")
	copyFile("./templates/index.js", "./"+dirName+"/"+"index.js")
}
