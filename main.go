package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("No additional arguments provided.")
		return
	}

	filePath := "sample.txt" // Change this to the path of your file
	file, err := os.ReadFile(filePath)

	fileData := string(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// fmt.Print(string(file))

	if strings.Contains(fileData, args[1]) {
		fmt.Printf("the term \"%s\" is present in file %s", args[1], filePath)
	} else {
		fmt.Printf("the term \"%s\" is not present in file %s", args[1], filePath)
	}
}
