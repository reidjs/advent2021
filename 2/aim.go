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
	aim := 0
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		dir := s[0]
		mag, _ := strconv.Atoi(s[1])
		if dir == "forward" {
			x += mag
			y += aim * mag
		} else if dir == "up" {
			aim -= mag
		} else if dir == "down" {
			aim += mag
		} else {
			fmt.Println("err")
		}
	}
	fmt.Println(x * y)
}
