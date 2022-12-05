package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
		return nil
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

func toInt(n string) int {
	num, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	return num
}
