package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sort_string(input string) string {
	s := strings.Split(input, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func decrypt(input []string, m map[int]string) map[int]string {
	// compare 1 and 4
	fmt.Println("m[1], m[4]", m[1], m[4])
	// group len(5) and len(6) strings
	// store them to process them later
	m5 := []string{}
	m6 := []string{}
	for _, el := range input {
		if len(el) == 5 {
			m5 = append(m5, sort_string(el))
		}
		if len(el) == 6 {
			m6 = append(m6, sort_string(el))
		}
	}
	fmt.Println("m5", m5)
	fmt.Println("m6", m6)
	return m
}

func main() {
	lines, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(lines)
	count := 0
	m := make(map[int]string)
	for scanner.Scan() {
		line := scanner.Text()
		t := strings.Split(line, "|")
		// u := strings.Split(t[1], " ")
		fmt.Println("t[0]", t[0])
		input := strings.Split(t[0], " ")
		fmt.Println("input", input)
		for _, el := range input {

			el = sort_string(el)
			if len(el) == 2 {
				m[1] = el
			}
			if len(el) == 3 {
				m[7] = el
			}
			if len(el) == 4 {
				m[4] = el
			}
			if len(el) == 7 {
				m[8] = el
			}
			decrypted_map := decrypt(strings.Split(t[0], " "), m)
			fmt.Println("decrypted_map", decrypted_map)
		}
		return

	}
	fmt.Println("count", count)
}
