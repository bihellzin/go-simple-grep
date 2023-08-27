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
	expressionLength := len(expression)
	filePath := args[2]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	red := "\033[31m"
	bold := "\033[1m"
	reset := "\033[0m"

	// fmt.Printf("%sTexto em vermelho e negrito%s\n", red+bold, reset)

	// coloredText := fmt.Sprintf("%sTexto vermelho%s e %sTexto verde%s", red+bold, reset, green, reset)
	// fmt.Println(coloredText)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, expression) {
			concat := ""
			n := 0
			positions := findAllOccurrences(line, expression)

			position := positions[n]

			for index, value := range line {
				if position == index {
					concat += red + bold
				}
				concat += string(value)

				if index >= position + expressionLength - 1 {
					concat += reset
					if n < len(positions) - 1 {
						n++
						position = positions[n]
					}
				}
			}
			fmt.Println(concat)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}
}
