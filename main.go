package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	Red = "\033[31m"
	Bold = "\033[1m"
	Reset = "\033[0m"
	Magenta = "\033[35m"
	Green = "\033[32m"
)

func findAllOccurrences(mainString, substring string) []int {
	positions := []int{}
	start := 0

	for {
		index := strings.Index(mainString[start:], substring)
		if index == -1 {
			break
		}
		positions = append(positions, start+index)
		start = start + index + len(substring)
	}

	return positions
}

func traverse(path, expression string) {
	dir, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading directory contents:", err)
		return
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			absolutePath, err := filepath.Abs(path+"/"+fileInfo.Name())
			if err != nil {
				fmt.Println("Error reading directory contents:", err)
				return
			}

			traverse(absolutePath, expression)
		} else {
			absolutePath, err := filepath.Abs(path+"/"+fileInfo.Name())
			if err != nil {
				fmt.Println("Error reading directory contents:", err)
				return
			}
			file, err := os.Open(absolutePath)
			readFileAndHighlight(file, expression, absolutePath)

			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer file.Close()
		}
	}
}

func highlightExpression(line, expression string, filePath ...string) {
	concat := ""
	expressionLength := len(expression)
	occurrenceIndex := 0
	positions := findAllOccurrences(line, expression)

	position := positions[occurrenceIndex]

	for index, value := range line {
		if position == index {
			concat += Red + Bold
		}
		concat += string(value)

		if index >= position + expressionLength - 1 {
			concat += Reset
			if occurrenceIndex < len(positions) - 1 {
				occurrenceIndex++
				position = positions[occurrenceIndex]
			}
		}
	}
	if (len(filePath) > 0) {
		fmt.Printf("%s%s%s%s:%s%s\n", Magenta, filePath[0], Reset, Green, Reset, concat)
		return
	}
	fmt.Println(concat)
}

func readFileAndHighlight(file io.Reader, expression string, filePath ...string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, expression) {
			if len(filePath) > 0 {
				highlightExpression(line, expression, filePath[0])
			} else {
				highlightExpression(line, expression)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("No additional arguments provided.")
		return
	}

	var expression string
	var path string
	var option string

	if len(args) == 3 {
		expression = args[1]
		path = args[2]
		file, err := os.Open(path)

		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}

		defer file.Close()
	}

	if len(args) == 4 {
		option = args[1]
		expression = args[2]
		path = args[3]
		file, err := os.Open(path)

		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}

		defer file.Close()
	}

	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Check if the path points to a file
	if info.Mode().IsRegular() {
		file, err := os.Open(path)
		readFileAndHighlight(file, expression)

		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
	} else if info.Mode().IsDir() && strings.Contains(option, "r") {
		traverse(path, expression)
		return
	} else {
		fmt.Println("Path is neither a file nor a directory.")
	}
}
