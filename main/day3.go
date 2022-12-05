package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func main3_2() {
	lines, err := readLines("resource/day3.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	sum := 0
	for i := 0; i < len(lines); i += 3 {
		str0 := lines[i]
		str1 := lines[i+1]
		str2 := lines[i+2]

		for _, s := range str0 {
			if strings.ContainsRune(str1, s) {
				if strings.ContainsRune(str2, s) {
					fmt.Println("string: ", strconv.QuoteRune(s), "number: ", int(s))
					if unicode.IsLower(s) {
						sum += int(s) - 96
					} else {
						sum += int(s) - 38
					}
					break
				}

			}
		}
	}
	fmt.Printf("Result is %d \n", sum)
}

func main3_1() {
	// lines, err := readLines("input.txt")
	lines, err := readLines("resource/day3.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	sum := 0
	for _, line := range lines {
		runee := []rune(line)
		str1 := string(runee[0 : len(runee)/2])
		str2 := string(runee[len(runee)/2:])

		fmt.Println("first string: ", str1)
		fmt.Println("second string: ", str2)

		for _, s := range str1 {
			if strings.ContainsRune(str2, s) {
				fmt.Println("string: ", strconv.QuoteRune(s), "number: ", int(s))
				if unicode.IsLower(s) {
					sum += int(s) - 96
				} else {
					sum += int(s) - 38
				}
				break
			}
		}
	}

	fmt.Printf("Result is %d \n", sum)

}
