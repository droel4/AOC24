package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	filePath := os.Args[1]

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	input := string(file)

	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)

	matches := re.FindAllStringSubmatch(input, -1)

	enabled := true
	sum := 0

	for _, match := range matches {
		token := match[0]

		switch {
		case token == "do()":
			enabled = true
		case token == "don't()":
			enabled = false
		default:
			if enabled {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				sum += x * y
			}
		}
	}

	fmt.Println("Sum of all multiplications:", sum)
}
