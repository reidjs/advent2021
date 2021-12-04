package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	x := 0
	y := 0
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// slice word from space
		s := strings.Fields(scanner.Text())
		dir := s[0]
		mag, _ := strconv.Atoi(s[1])
		if dir == "forward" {
			x += mag
		} else if dir == "up" {
			y -= mag
		} else if dir == "down" {
			y += mag
		} else {
			fmt.Println("err")
		}
	}
	fmt.Println(x * y)

}
