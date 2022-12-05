package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
[P]     [C]         [M]
[D]     [P] [B]     [V] [S]
[Q] [V] [R] [V]     [G] [B]
[R] [W] [G] [J]     [T] [M]     [V]
[V] [Q] [Q] [F] [C] [N] [V]     [W]
[B] [Z] [Z] [H] [L] [P] [L] [J] [N]
[H] [D] [L] [D] [W] [R] [R] [P] [C]
[F] [L] [H] [R] [Z] [J] [J] [D] [D]
 1   2   3   4   5   6   7   8   9
 */

func main() {
	lines := readLines("resource/day5.txt")

	queue1 := make([]int, 0)
// Push to the queue
queue = append(queue, 1)
// Top (just get next element, don't remove it)
x = queue[0]
// Discard top element
queue = queue[1:]
// Is empty ?
if len(queue) == 0 {
    fmt.Println("Queue is empty !")
}


	for _, line := range lines {
		slice := strings.Split(line, ",")
		fmt.Println("slice is: ", slice)
		left := strings.Split(slice[0], "-")
		fmt.Println("left is: ", left)
		right := strings.Split(slice[1], "-")
		fmt.Println("right is: ", right)


	}
	fmt.Printf("Result is %d \n", result)
}