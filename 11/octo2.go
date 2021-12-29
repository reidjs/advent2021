package main

import "fmt"

const SIZE = 10

var grid = [SIZE][SIZE]int{
	{5, 4, 2, 1, 4, 5, 1, 7, 4, 1},
	{3, 8, 7, 7, 3, 2, 1, 5, 6, 8},
	{7, 5, 8, 3, 2, 7, 3, 8, 6, 4},
	{3, 4, 5, 1, 7, 1, 7, 7, 7, 8},
	{2, 6, 5, 1, 6, 1, 5, 1, 5, 6},
	{6, 3, 7, 7, 1, 6, 7, 5, 2, 6},
	{5, 1, 8, 2, 8, 5, 2, 8, 3, 1},
	{4, 7, 6, 6, 8, 5, 6, 6, 7, 6},
	{3, 4, 3, 7, 1, 8, 7, 5, 8, 3},
	{3, 6, 3, 3, 3, 7, 1, 5, 8, 6},
}

func increase_if_possible(row int, col int) {
	val := grid[row][col]
	if val != 0 {
		grid[row][col]++
	}
}

func left(row int, col int) {
	if col-1 >= 0 {
		increase_if_possible(row, col-1)
	}
}
func right(row int, col int) {
	if col+1 < SIZE {
		increase_if_possible(row, col+1)
	}
}

func up(row int, col int) {
	if row-1 >= 0 {
		increase_if_possible(row-1, col)
	}
}
func down(row int, col int) {
	if row+1 < SIZE {
		increase_if_possible(row+1, col)
	}
}
func left_up(row int, col int) {
	if col-1 >= 0 && row-1 >= 0 {
		increase_if_possible(row-1, col-1)
	}
}
func left_down(row int, col int) {
	if col-1 >= 0 && row+1 < SIZE {
		increase_if_possible(row+1, col-1)
	}
}
func right_up(row int, col int) {
	if col+1 < SIZE && row-1 >= 0 {
		increase_if_possible(row-1, col+1)
	}
}
func right_down(row int, col int) {
	if col+1 < SIZE && row+1 < SIZE {
		increase_if_possible(row+1, col+1)
	}
}

func flash(row int, col int) {
	left(row, col)
	right(row, col)
	up(row, col)
	down(row, col)
	left_up(row, col)
	right_up(row, col)
	left_down(row, col)
	right_down(row, col)
}

func check_if_all_flashed() {
	result := true
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] != 0 {
				result = false
			}
		}
	}
	if result {
		fmt.Println("done")
	}
}

func process_flashes() {
	check_if_all_flashed()
	flash_occurred := false
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] > 9 {
				grid[row][col] = 0
				flash(row, col)
				flash_occurred = true
			}
		}
	}
	if flash_occurred {
		process_flashes()
	}
}

func main() {
	i := 0
	for i < 1000 {
		for row := range grid {
			for col := range grid[row] {
				grid[row][col]++
			}
		}
		fmt.Println("i", i)
		process_flashes()
		i++
	}

}
