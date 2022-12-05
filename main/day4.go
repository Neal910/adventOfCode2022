package main

import (
	"fmt"
	"strings"
)

func main4_2() {
	lines := readLines("resource/day4.txt")

	result := 0
	for _, line := range lines {
		slice := strings.Split(line, ",")
		fmt.Println("slice is: ", slice)
		left := strings.Split(slice[0], "-")
		fmt.Println("left is: ", left)
		right := strings.Split(slice[1], "-")
		fmt.Println("right is: ", right)

		if (toInt(left[0]) <= toInt(right[0]) && toInt(left[1]) >= toInt(right[0])) || (toInt(right[0]) <= toInt(left[0]) && toInt(left[0]) <= toInt(right[1])) {
			result++
		}

	}
	fmt.Printf("Result is %d \n", result)
}
