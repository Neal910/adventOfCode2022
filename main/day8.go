package main

import (
	"fmt"
	"sort"
	"strings"
)

func main8_1_2() {
	lines := readLines("./resource/day8.txt")
	// lines := readLines("./resource/test.txt")

	edgeTreeSubTotal := 2*len(lines) + 2*(len(lines[0])-2)

	interiorTreeSubTotal, maxDistance := findInteriorVisibleTree(lines)

	fmt.Println("All visible tree number: ", edgeTreeSubTotal+interiorTreeSubTotal)
	fmt.Println("Highest scenic score: ", maxDistance)

}

func findInteriorVisibleTree(trees []string) (int, int) {
	sum := 0
	var distances []int
	twoDimensionTree := convertTo2DArray(trees)
	for i := 1; i < len(twoDimensionTree)-1; i++ {
		for j := 1; j < len(twoDimensionTree[0])-1; j++ {
			boolean, dis := isVisible(i, j, twoDimensionTree[i][j], twoDimensionTree)
			if boolean {
				sum++
				distances = append(distances, dis)
			}
		}
	}
	sort.Ints(distances)
	return sum, distances[len(distances)-1]
}

func isVisible(row, col, height int, trees [][]int) (bool, int) {

	left, right, up, down := true, true, true, true
	leftDis, rightDis, upDis, downDis := 0, 0, 0, 0
	distance := 0

	for i := col - 1; i >= 0; i-- {
		if trees[row][i] >= height {
			left = false
			leftDis = col - i
			break
		} else {
			leftDis = col
		}
	}
	for i := col + 1; i < len(trees); i++ {
		if trees[row][i] >= height {
			right = right && false
			rightDis = i - col
			break
		} else {
			rightDis = len(trees) - 1 - col
		}
	}
	for j := row + 1; j < len(trees[0]); j++ {
		if trees[j][col] >= height {
			down = down && false
			downDis = j - row
			break
		} else {
			downDis = len(trees) - 1 - row
		}
	}
	for j := row - 1; j >= 0; j-- {
		if trees[j][col] >= height {
			up = up && false
			upDis = row - j
			break
		} else {
			upDis = row
		}
	}
	if left || right || up || down {
		distance = leftDis * rightDis * upDis * downDis
	}

	return left || right || up || down, distance
}

func convertTo2DArray(arrays []string) [][]int {
	rowNumber := len(arrays)
	colNumber := len(strings.Split(arrays[0], ""))

	twoDArr := make([][]int, rowNumber, colNumber)

	for i, arr := range arrays {
		t := strings.Split(arr, "")
		twoDArr[i] = append(twoDArr[i], toArrayInt(t)...)
	}
	return twoDArr
}
