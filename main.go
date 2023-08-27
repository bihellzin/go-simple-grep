package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("No additional arguments provided.")
		return
	}

	expression := args[1]
	filePath := args[2]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// red := "\033[31m"
	// bold := "\033[1m"
	// reset := "\033[0m"
	// green := "\033[32m"

	// fmt.Printf("%sTexto em vermelho e negrito%s\n", red+bold, reset)

	// coloredText := fmt.Sprintf("%sTexto vermelho%s e %sTexto verde%s", red+bold, reset, green, reset)
	// fmt.Println(coloredText)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, expression) {
			fmt.Println(line)
			// positions := findAllOccurrences(line, expression)
			// fmt.Println(positions)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}
}
