package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkRowColForBingo(board [][]int, row int, col int) bool {
	size := len(board[0])
	rowWinner := true
	colWinner := true
	for _, cell := range board[row%size] {
		if cell != -1 {
			rowWinner = false
		}
	}
	for i := range board {
		cell := board[i][col]
		if cell != -1 {
			colWinner = false
		}
	}
	return colWinner || rowWinner
}

func sumOfPositiveValues(board [][]int) int {
	sum := 0
	for _, col := range board {
		for _, cell := range col {
			if cell > 0 {
				sum += cell
			}
		}
	}
	return sum
}

/**
* returns [turns, calculatedBoardValue] turns it takes fort the board to get bingo (vertical or horizontal)
* calculated board value is the sum of all unmarked numbers on the board, multiplied by the number that was called when the board won
* returns [-1, -1] if does not win
 */
func turnsToWin(board [][]int, numbers []int) []int {
	for turns, n := range numbers {
		for i, v := range board {
			for j, cell := range v {
				if cell == n {
					board[i][j] = -1
				}
				win := checkRowColForBingo(board, i, j)
				if win {
					// fmt.Println("turns", turns)
					// fmt.Println("n", n)
					return []int{turns, n * sumOfPositiveValues(board)}
				}
			}
		}
	}

	// if we have reached the end of numbers, this is a nonwinning board, return -1, -1
	return []int{-1, -1}
}

func main() {
	// board := [][]int{
	// 	{14, 21, 17, 24, 4},
	// 	{10, 16, 15, 9, 19},
	// 	{18, 8, 23, 26, 20},
	// 	{22, 11, 13, 6, 5},
	// 	{2, 0, 12, 3, 7},
	// }
	numbers := []int{87, 7, 82, 21, 47, 88, 12, 71, 24, 35, 10, 90, 4, 97, 30, 55, 36, 74, 19, 50, 23, 46, 13, 44, 69, 27, 2, 0, 37, 33, 99, 49, 77, 15, 89, 98, 31, 51, 22, 96, 73, 94, 95, 18, 52, 78, 32, 83, 85, 54, 75, 84, 59, 25, 76, 45, 20, 48, 9, 28, 39, 70, 63, 56, 5, 68, 61, 26, 58, 92, 67, 53, 43, 62, 17, 81, 80, 66, 91, 93, 41, 64, 14, 8, 57, 38, 34, 16, 42, 11, 86, 72, 40, 65, 79, 6, 3, 29, 60, 1}
	fmt.Println("numbers", numbers)
	// numbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	// read in boards
	// TODO: read in all numbers and store correctly
	// TODO: read in all boards
	rawBoards, _ := os.Open("boards.txt")
	scanner := bufio.NewScanner(rawBoards)
	// board := [][]int
	board := [][]int{}
	// turns, value
	currentWinner := []int{10000, -1}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			fmt.Println("Checking winner")
			// TODO: check for winner
			x := turnsToWin(board, numbers)
			fmt.Println("x", x)
			turns := x[0]
			if turns < currentWinner[0] {
				currentWinner = x
			}
			board = [][]int{}
		} else {
			nums := strings.Fields(scanner.Text())
			var foo []int
			for _, n := range nums {
				m, _ := strconv.Atoi(n)
				foo = append(foo, m)
			}
			board = append(board, foo)
			// fmt.Println("b", board)
			// board.append()
		}
	}
	fmt.Println("currentWinner", currentWinner)
	// test := [][]int{
	// 	{-1, 2, -1},
	// 	{0, 0, 0},
	// }
	// y := checkRowColForBingo(test, 0, 0)
	// fmt.Println("y", y)

	// iterate through each board and pass in numbers, seeing how many moves until it wins (or -1 for no win)
}
