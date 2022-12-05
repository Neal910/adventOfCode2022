package main

import (
	"fmt"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func main5_2() {
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
	stack1 := stack.New()
	stack1.Push('F')
	stack1.Push('H')
	stack1.Push('B')
	stack1.Push('V')
	stack1.Push('R')
	stack1.Push('Q')
	stack1.Push('D')
	stack1.Push('P')

	stack2 := stack.New()
	stack2.Push('L')
	stack2.Push('D')
	stack2.Push('Z')
	stack2.Push('Q')
	stack2.Push('W')
	stack2.Push('V')

	stack3 := stack.New()
	stack3.Push('H')
	stack3.Push('L')
	stack3.Push('Z')
	stack3.Push('Q')
	stack3.Push('G')
	stack3.Push('R')
	stack3.Push('P')
	stack3.Push('C')

	stack4 := stack.New()
	stack4.Push('R')
	stack4.Push('D')
	stack4.Push('H')
	stack4.Push('F')
	stack4.Push('J')
	stack4.Push('V')
	stack4.Push('B')

	stack5 := stack.New()
	stack5.Push('Z')
	stack5.Push('W')
	stack5.Push('L')
	stack5.Push('C')

	stack6 := stack.New()
	stack6.Push('J')
	stack6.Push('R')
	stack6.Push('P')
	stack6.Push('N')
	stack6.Push('T')
	stack6.Push('G')
	stack6.Push('V')
	stack6.Push('M')

	stack7 := stack.New()
	stack7.Push('J')
	stack7.Push('R')
	stack7.Push('L')
	stack7.Push('V')
	stack7.Push('M')
	stack7.Push('B')
	stack7.Push('S')

	stack8 := stack.New()
	stack8.Push('D')
	stack8.Push('P')
	stack8.Push('J')

	stack9 := stack.New()
	stack9.Push('D')
	stack9.Push('C')
	stack9.Push('N')
	stack9.Push('W')
	stack9.Push('V')

	lines := readLines("resource/day5.txt")

	for _, line := range lines {
		slice := strings.Split(line, " ")
		var st1 *stack.Stack
		var st2 *stack.Stack

		switch slice[3] {
		case "1":
			st1 = stack1
		case "2":
			st1 = stack2
		case "3":
			st1 = stack3
		case "4":
			st1 = stack4
		case "5":
			st1 = stack5
		case "6":
			st1 = stack6
		case "7":
			st1 = stack7
		case "8":
			st1 = stack8
		case "9":
			st1 = stack9
		}

		switch slice[5] {
		case "1":
			st2 = stack1
		case "2":
			st2 = stack2
		case "3":
			st2 = stack3
		case "4":
			st2 = stack4
		case "5":
			st2 = stack5
		case "6":
			st2 = stack6
		case "7":
			st2 = stack7
		case "8":
			st2 = stack8
		case "9":
			st2 = stack9
		}

		move(toInt(slice[1]), st1, st2)

	}
	fmt.Println("Result is: ", stack1.Peek(), stack2.Peek(), stack3.Peek(), stack4.Peek(), stack5.Peek(), stack6.Peek(), stack7.Peek(), stack8.Peek(), stack9.Peek())

	fmt.Println(string(86), string(72), string(74), string(68), string(68), string(67), string(87), string(82), string(68))
}

func move(num int, st1 *stack.Stack, st2 *stack.Stack) {
	var tmp stack.Stack
	for i := 0; i < num; i++ {
		tmp.Push(st1.Pop())
	}
	for i := 0; i < num; i++ {
		st2.Push(tmp.Pop())
	}
}
