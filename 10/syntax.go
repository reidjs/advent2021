package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	m := map[string]string{
		"{": "}",
		"<": ">",
		"[": "]",
		"(": ")",
	}
	m2 := map[string]string{
		"}": "{",
		">": "<",
		"]": "[",
		")": "(",
	}
	m3 := map[string]int{
		">": 25137,
		"}": 1197,
		"]": 57,
		")": 3,
	}
	score := 0
	// input := []string{
	// 	"[({(<(())[]>[[{[]{<()<>>",
	// 	"[(()[<>])]({[<{<<[]>>(",
	// 	"{([(<{}[<>[]}>{[]{[(<()>",
	// 	"[[<[([]))<([[{}[[()]]]",
	// 	"[{[{({}]{}}([{[{{{}}([]",
	// 	"[<(<(<(<{}))><([]([]()",
	// 	"<{([([[(<>()){}]>(<<{{",
	// }
	// fmt.Println("input", input)
	rawInput, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(rawInput)
	for scanner.Scan() {
		line := scanner.Text()
		stack := []string{}
		for i := 0; i < len(line); i++ {
			el := string(line[i])
			if _, ok := m[el]; ok {
				stack = append(stack, el)
			} else {
				// check for validity
				peek := stack[len(stack)-1]
				if peek != m2[el] {
					// invalid
					// fmt.Println("invalid el", el)
					score += m3[el]
					break
				}
				pop := len(stack) - 1
				stack = stack[:pop]
			}
			// fmt.Println("stack", stack)
		}
	}
	fmt.Println("score", score)
}
