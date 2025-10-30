package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var leftNumbers []int
	var rightNumbers []int

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Split by whitespace (handles multiple spaces or tabs)
		fields := strings.Fields(line)
		if len(fields) != 2 {
			log.Printf("Skipping malformed line %d: %s\n", lineNumber, line)
			continue
		}

		left, err1 := strconv.Atoi(fields[0])
		right, err2 := strconv.Atoi(fields[1])

		if err1 != nil || err2 != nil {
			log.Printf("Skipping line %d (invalid number): %s\n", lineNumber, line)
			continue
		}

		leftNumbers = append(leftNumbers, left)
		rightNumbers = append(rightNumbers, right)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	var diffTotal int = 0

	for i := 0; i < len(leftNumbers); i++ {
		dif := leftNumbers[i] - rightNumbers[i]
		diffTotal += abs(dif)
	}

	var multTotal int = 0
	for i := 0; i < len(leftNumbers); i++ {
		count := 0
		for j := 0; j < len(rightNumbers); j++ {
			if leftNumbers[i] == rightNumbers[j] {
				count++
			}
		}
		multTotal += count * leftNumbers[i]
	}

	fmt.Println("Total Difference:", diffTotal)
	fmt.Println("Total Multiply:", multTotal)
}
