package main

import (
	"fmt"
	"strings"
)

// A for Rock, B for Paper, and C for Scissors.
// X for Rock, Y for Paper, and Z for Scissors
// 1 for Rock, 2 for Paper, and 3 for Scissors
// 0 if you lost, 3 if the round was a draw, and 6 if you won

// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win.
func main2_2() {
	lines := readLines("input.txt")

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
