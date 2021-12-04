package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner((file))
	currentSum := -1
	prevSum := -1
	prevValue := -1
	prev2Value := -1
	increaseCount := 0
	var index int = 0
	for scanner.Scan() {
		currentValue, _ := strconv.Atoi(scanner.Text())
		if index == 0 {
			prev2Value = currentValue
			index++
			continue
		}
		if index == 1 {
			prevValue = currentValue
			index++
			continue
		}
		if index == 2 {
			prevSum = currentValue + prevValue + prev2Value
			prev2Value = prevValue
			prevValue = currentValue
			index++
			continue
		}
		currentSum = currentValue + prevValue + prev2Value
		fmt.Printf("currentSum: %v\n", currentSum)
		fmt.Printf("prevSum: %v\n", prevSum)
		if currentSum > prevSum && prevSum >= 0 {
			increaseCount++
			fmt.Println("Increased sum")
		}
		prev2Value = prevValue
		prevValue = currentValue
		prevSum = currentSum
		index++
	}
	fmt.Println(increaseCount)
}
