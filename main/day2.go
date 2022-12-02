package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// A for Rock, B for Paper, and C for Scissors.
// X for Rock, Y for Paper, and Z for Scissors
// 1 for Rock, 2 for Paper, and 3 for Scissors
// 0 if you lost, 3 if the round was a draw, and 6 if you won

// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win.
func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	tmp := 0
	sum := 0

	for _, line := range lines {
		slic := strings.Split(line, " ")
		switch slic[0] {
		case "A":
			switch slic[1] {
			case "X":
				tmp = 3 + 0
			case "Y":
				tmp = 1 + 3
			case "Z":
				tmp = 2 + 6
			}
		case "B":
			switch slic[1] {
			case "X":
				tmp = 1 + 0
			case "Y":
				tmp = 2 + 3
			case "Z":
				tmp = 3 + 6
			}
		case "C":
			switch slic[1] {
			case "X":
				tmp = 2 + 0
			case "Y":
				tmp = 3 + 3
			case "Z":
				tmp = 1 + 6
			}
		}
		sum += tmp
	}

	fmt.Printf("Result is %d \n", sum)

}
