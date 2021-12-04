// life support = oxygen generator rating * co2 scrubber rating
// filtering out values until only one remains
// start w/ full list of binary numbers, consider just the first bit
// keep only numbers selected by the bit criteria for the type of rating, discard nums that do not match
// if only one number left, stop, this is the rating value
// otherwise, repeat process, considering next bit to right

// bit critera
// oxygen rating: determine most common value in current bit position
// and keep only numbers with that bit in that position
// if 0 and 1 are equally common, keep values w/ a 1 in the position being considered

// co2 rating: determine least common value
// keep only numbers with that bit in that position
// if 0 and 1 equally common, use 0
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getWinningCount(arr []string, bitIdx int) int {
	var counts = []int{0, 0}

	for _, c := range arr {
		val := string(c[bitIdx])
		if val == "0" {
			counts[0]++
		} else if val == "1" {
			counts[1]++
		} else {
			fmt.Println("err!!!")
		}
	}

	result := -1
	if counts[0] > counts[1] {
		result = 0
	} else if counts[1] > counts[0] {
		result = 1
	}

	return result
}

func applyFilter(arr []string, bitIdx int, value string) (result []string) {
	for _, c := range arr {
		if string(c[bitIdx]) == value {
			result = append(result, string(c))
		}
	}
	return
}

func oxy(arr []string, bitIdx int) string {
	if len(arr) == 1 {
		return arr[0]
	}
	// get the winning count of the column
	winningCount := getWinningCount(arr, bitIdx)
	// filteredArr := []string

	if winningCount == 0 {
		// filter out binary numbers that don't start with winning value
		arr = applyFilter(arr, bitIdx, "0")
	} else if winningCount == 1 {
		arr = applyFilter(arr, bitIdx, "1")
	} else {
		// default to 1
		arr = applyFilter(arr, bitIdx, "1")
	}
	return oxy(arr, bitIdx+1)
}
func co2(arr []string, bitIdx int) string {
	if len(arr) == 1 {
		return arr[0]
	}
	// get the winning count of the column
	winningCount := getWinningCount(arr, bitIdx)
	// filteredArr := []string

	if winningCount == 0 {
		// filter out binary numbers that don't start with winning value
		arr = applyFilter(arr, bitIdx, "1")
	} else if winningCount == 1 {
		arr = applyFilter(arr, bitIdx, "0")
	} else {
		// default to 1
		arr = applyFilter(arr, bitIdx, "0")
	}
	fmt.Println("arr", arr)
	return co2(arr, bitIdx+1)
}

func main() {
	var s []string
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	oxygen := oxy(s, 0)
	co2scrub := co2(s, 0)
	fmt.Println("oxygen", oxygen)

	fmt.Println("co2scrub", co2scrub)
	a, _ := strconv.ParseInt(oxygen, 2, 64)
	b, _ := strconv.ParseInt(co2scrub, 2, 64)
	fmt.Println(a * b)
}
