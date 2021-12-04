package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// power consumption = gamma_rate * epsilon_rate
// gamma rate: consider only the first bit of each binary number -> 00100 --> 0, 111010 --> 1, so that becomes 01
// use that to build up a binary number for gamma rate(the most common val)
// epsilon rate uses the least common bit from each position

func main() {
	// use array of arrays to track counts [[zeros, ones], [zeros, ones]]
	var counts [12][2]int
	for i := range counts {
		counts[i][0] = 0
		counts[i][1] = 0
	}
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		for i, c := range s {
			val := string(c)
			if val == "0" {
				counts[i][0]++
			} else if val == "1" {
				counts[i][1]++
			} else {
				fmt.Println("err!!!")
			}
		}
	}
	gamma := ""
	epsilon := ""
	// loop thru counts to build up binary vals
	for i := range counts {
		if counts[i][0] > counts[i][1] {
			gamma += "0"
			epsilon += "1"
		} else if counts[i][1] > counts[i][0] {
			gamma += "1"
			epsilon += "0"
		} else {
			fmt.Println("err2")
		}
	}
	a, _ := strconv.ParseInt(gamma, 2, 64)
	b, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(a * b)

}
