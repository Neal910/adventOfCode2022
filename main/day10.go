package main

import (
	"fmt"
	"strings"
)

func main10_1() {
	lines := readLines("resource/day10.txt")

	cycles := 0
	value := 1
	total := 0

	for _, line := range lines {
		arr := strings.Split(line, " ")

		if arr[0] == "noop" {
			cycles++
			isCycle, val := cycleCheck(cycles, value)
			if isCycle {
				total += val
			}
		} else if arr[0] == "addx" {
			cycles++
			isCycle, val := cycleCheck(cycles, value)
			if isCycle {
				total += val
			}
			cycles++
			isCycle, val = cycleCheck(cycles, value)
			if isCycle {
				total += val
			}
			value += toInt(arr[1])
		}
	}
	fmt.Println("Sum of these six signal strengths: ", total)
}

func main10_2() {
	lines := readLines("resource/day10.txt")

	cycles := 0
	sprite_position := 1

	for _, line := range lines {
		arr := strings.Split(line, " ")

		if arr[0] == "noop" {
			cycles++
			cycles = printSignal(cycles, sprite_position)
		} else if arr[0] == "addx" {
			cycles++
			cycles = printSignal(cycles, sprite_position)
			cycles++
			cycles = printSignal(cycles, sprite_position)
			sprite_position += toInt(arr[1])
		}
	}
}

func printSignal(cyc, pos int) int {
	if pos <= cyc && cyc <= pos+2 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}

	if cyc == 40 {
		fmt.Printf("\n")
		return cyc - 40
	}
	return cyc
}

func cycleCheck(cyc, val int) (bool, int) {
	boolean := false
	strength := 0
	//20th, 60th, 100th, 140th, 180th, and 220th cycles
	switch cyc {
	case 20:
		boolean = true
		strength = val * 20
	case 60:
		boolean = true
		strength = val * 60
	case 100:
		boolean = true
		strength = val * 100
	case 140:
		boolean = true
		strength = val * 140
	case 180:
		boolean = true
		strength = val * 180
	case 220:
		boolean = true
		strength = val * 220
	}
	return boolean, strength
}
