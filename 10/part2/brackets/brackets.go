package brackets

import (
	"bufio"
	"fmt"
	"os"
)

func CalculateScore(stack []string) int {
	m3 := map[string]int{
		">": 4,
		"}": 3,
		"]": 2,
		")": 1,
	}
	score := 0
	for _, s := range stack {
		score = (score * 5) + m3[s]
	}
	return score
}

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

	scores := []int{}
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
					// score += m3[el]
					break
				}
				pop := len(stack) - 1
				stack = stack[:pop]
			}
			fmt.Println("stack", stack)
			if len(stack) != 0 {
				// incomplete, calc score
				scores = append(scores, CalculateScore(stack))
			}
		}
	}
	// fmt.Println("score", score)
}
