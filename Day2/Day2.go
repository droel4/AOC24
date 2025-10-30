package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func arrayPasses(array []int) bool {
	if len(array) < 1 {
		return false
	}
	increasing := false
	decreasing := false
	prevNum := array[0]
	for i := 1; i < len(array); i++ {
		currentNum := array[i]
		if currentNum == prevNum {
			return false
		}
		if currentNum < prevNum {
			decreasing = true
		}
		if currentNum > prevNum {
			increasing = true
		}
		if abs(currentNum-prevNum) > 3 {
			return false
		}
		prevNum = currentNum
	}
	if increasing && decreasing {
		return false
	}
	return true
}

// removeOne generates all slices with one element removed from the original.
func removeOneElement(nums []int) [][]int {
	result := [][]int{}

	for i := range nums {
		// Create a new slice that skips the element at index i
		newSlice := append([]int{}, nums[:i]...)   // elements before i
		newSlice = append(newSlice, nums[i+1:]...) // elements after i
		result = append(result, newSlice)
	}

	return result
}

func arrayPassesErrorRemoved(array []int) bool {
	if arrayPasses(array) {
		return true
	}
	arrays := removeOneElement(array)
	for i := 0; i < len(arrays); i++ {
		if arrayPasses(arrays[i]) {
			return true
		}
	}
	return false
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

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	var numbers [][]int

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Split by whitespace (handles multiple spaces or tabs)
		fields := strings.Fields(line)

		var list []int

		for i := 0; i < len(fields); i++ {
			val, err := strconv.Atoi(fields[i])
			if err != nil {
				log.Printf("Skipping line %d (invalid number): %s\n", lineNumber, line)
				continue
			}
			list = append(list, val)
		}

		numbers = append(numbers, list)
	}

	passCount := 0
	almostPassCount := 0
	for i := 0; i < len(numbers); i++ {
		if arrayPasses(numbers[i]) {
			passCount++
		}
		if arrayPassesErrorRemoved(numbers[i]) {
			almostPassCount++
		}
	}

	fmt.Println("Number of passes:", passCount)
	fmt.Println("Number of almost passes:", almostPassCount)
}
