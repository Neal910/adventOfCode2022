package main

import "fmt"

func main() {
	lines := readLines("./resource/day6.txt")

	line := lines[0]
	for i := 0; i < len(line); i++ {
		str := line[i : i+14]
		arr := []rune(str)
		if !isDuplicate(arr) {
			fmt.Println("Result is: ", i+14)
			return
		}
	}
	fmt.Println("No result")
}
