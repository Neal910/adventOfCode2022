package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return lines
}

// toInt converts string to int
func toInt(n string) int {
	num, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return num
}

// isDuplicate check if the rune array has duplicate element
func isDuplicate(arr []rune) bool {
	visited := make(map[rune]bool, 0)

	for _, r := range arr {
		if visited[r] == true {
			return true
		} else {
			visited[r] = true
		}
	}
	return false
}
