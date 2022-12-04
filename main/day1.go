package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main1_2() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	tmp := 0
	arr := []int{}
	l := 0

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) > 0 {
			i, err := strconv.Atoi(lines[i])
			if err != nil {
				log.Fatal("string converting to int error")
			}
			tmp += i
		} else {
			arr = append(arr, tmp)
			tmp = 0
		}
	}
	arr = append(arr, tmp)
	sort.Ints(arr)
	l = len(arr)
	max := arr[l-1] + arr[l-2] + arr[l-3]

	fmt.Printf("Result is %d \n", max)

}
