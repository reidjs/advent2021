package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MAX_VAL = 10

func left(points [][]int, row int, col int) int {
	if col-1 < 0 {
		return MAX_VAL
	}
	return points[row][col-1]
}
func right(points [][]int, row int, col int) int {
	if col+1 >= len(points[row]) {
		return MAX_VAL
	}
	return points[row][col+1]
}
func up(points [][]int, row int, col int) int {
	if row-1 < 0 {
		return MAX_VAL
	}
	return points[row-1][col]
}
func down(points [][]int, row int, col int) int {
	if row+1 >= len(points) {
		return MAX_VAL
	}
	return points[row+1][col]
}

func isLowestPoint(points [][]int, row int, col int) bool {
	p := points[row][col]
	left := p < left(points, row, col)
	right := p < right(points, row, col)
	up := p < up(points, row, col)
	down := p < down(points, row, col)
	if left && right && up && down {
		return true
	}
	return false
}

func main() {
	// input := []string{
	// 	"2199943210",
	// 	"3987894921",
	// 	"9856789892",
	// 	"8767896789",
	// 	"9899965678",
	// }
	input := []string{}
	rawInput, _ := os.Open("input.txt")
	scanner := bufio.NewScanner((rawInput))
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	points := [][]int{}

	for i := 0; i < len(input); i++ {
		row := []int{}
		for j := 0; j < len(input[i]); j++ {
			cell, _ := strconv.Atoi(string(input[i][j]))
			row = append(row, cell)
		}

		points = append(points, row)
	}
	sum := 0
	for row := 0; row < len(points); row++ {
		for col := 0; col < len(points[row]); col++ {
			p := points[row][col]
			if isLowestPoint(points, row, col) {
				sum += 1 + p
			}
		}
	}
	fmt.Println("sum", sum)
	// fmt.Println("points", points)
	// fmt.Println("left(points[0][0])", left(points, 0, 0))
	// fmt.Println("right(points, 0, 0)", right(points, 0, 0))
	// fmt.Println("left(points, 0, ", left(points, 0, 9))
	// fmt.Println("right(points, 0, 9)", right(points, 0, 9))
	// fmt.Println("up(points, 0, 0)", up(points, 0, 0))
	// fmt.Println("down(points, 0, 0)", down(points, 0, 0))
	// fmt.Println("up(points, 0, 9)", up(points, 0, 9))
	// fmt.Println("down(points, 0, 9)", down(points, 0, 9))
}
