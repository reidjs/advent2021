package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	lastLine := -1
	count := 0
	for scanner.Scan() {
		currentLine, _ := strconv.Atoi(scanner.Text())
		if currentLine > lastLine {
			count++
		}
		lastLine = currentLine
	}
	fmt.Println(count - 1)
}
