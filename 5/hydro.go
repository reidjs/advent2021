package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// create hash table, initializes to all 0s
	// todo: this should be a singlr table, string: int
	// where string is x,y like 5,8 and int is the count of vectors crossing it
	// var mx map[int]int
	// var my map[int]int
	// mx = make(map[int]int)
	// my = make(map[int]int)
	var m map[string]int
	m = make(map[string]int)
	// m[123] = 1
	// fmt.Println("m[123]", m[124])
	// for each input, write the line of values
	lines, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(lines)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		p1 := strings.Split(line[0], ",")
		p2 := strings.Split(line[len(line)-1], ",")
		x1, _ := strconv.Atoi(p1[0])
		y1, _ := strconv.Atoi(p1[1])
		x2, _ := strconv.Atoi(p2[0])
		y2, _ := strconv.Atoi(p2[1])

		if x1 == x2 {
			if y2 >= y1 {
				for i := y1; i <= y2; i++ {
					strx := strconv.Itoa(x1)
					stry := strconv.Itoa(i)
					str := strx + "," + stry
					m[str]++
				}
			} else {
				for i := y2; i <= y1; i++ {
					strx := strconv.Itoa(x1)
					stry := strconv.Itoa(i)
					str := strx + "," + stry
					m[str]++
				}
			}
		} else if y1 == y2 {
			if x2 >= x1 {
				for i := x1; i <= x2; i++ {
					strx := strconv.Itoa(i)
					stry := strconv.Itoa(y1)
					str := strx + "," + stry
					m[str]++
				}
			} else {
				for i := x2; i <= x1; i++ {
					strx := strconv.Itoa(i)
					stry := strconv.Itoa(y1)
					str := strx + "," + stry
					m[str]++
				}
			}
		} else {
			// TODO: Not currently working. How to handle diagonals?
			// We can assume all right angles:
			if math.Abs(float64(x1-x2)) != math.Abs(float64(y1-y2)) {
				fmt.Println("ERROR")
			}
			// angle := math.Atan(math.Abs(float64(x1-x2)) / math.Abs(float64(y1-y2)))
			// they are all at 45 degree angles
			// find points between the vectors
			if y2 > y1 && x2 > x1 {
				// up one over one
				// 195,463 -> 658,926
				// fmt.Println("1")
				for i := 0; i <= y2-y1; i++ {
					strx := strconv.Itoa(x1 + i)
					stry := strconv.Itoa(y1 + i)
					str := strx + "," + stry
					m[str]++
				}
			} else if y2 > y1 && x2 < x1 {
				for i := 0; i <= y2-y1; i++ {
					strx := strconv.Itoa(x2 + i)
					stry := strconv.Itoa(y1 + i)
					str := strx + "," + stry
					m[str]++
				}
			} else if y2 < y1 && x2 > x1 {
				for i := 0; i <= y1-y2; i++ {
					strx := strconv.Itoa(x1 + i)
					stry := strconv.Itoa(y2 + i)
					str := strx + "," + stry
					m[str]++
				}
			} else if y2 < y1 && x2 < x1 {
				for i := 0; i <= y1-y2; i++ {
					strx := strconv.Itoa(x2 + i)
					stry := strconv.Itoa(y2 + i)
					str := strx + "," + stry
					m[str]++
				}
			} else {
				fmt.Println("ERROR")
			}
		}
		// fmt.Println("m", m)
	}
	count := 0
	for _, v := range m {
		if v > 1 {
			count++
		}
	}
	fmt.Println("count", count)
}
