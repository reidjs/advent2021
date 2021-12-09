package main

import "fmt"

func getCounts(fish []int) [8]int {
	counts := [8]int{}
	for i := 0; i < len(fish); i++ {
		counts[fish[i]]++
	}
	return counts
}

func main() {
	// fish := []int{1, 1, 5, 2, 1, 1, 5, 5, 3, 1, 1, 1, 1, 1, 1, 3, 4, 5, 2, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 1, 5, 4, 5, 1, 5, 3, 1, 3, 2, 1, 1, 1, 1, 2, 4, 1, 5, 1, 1, 1, 4, 4, 1, 1, 1, 1, 1, 1, 3, 4, 5, 1, 1, 2, 1, 1, 5, 1, 1, 4, 1, 4, 4, 2, 4, 4, 2, 2, 1, 2, 3, 1, 1, 2, 5, 3, 1, 1, 1, 4, 1, 2, 2, 1, 4, 1, 1, 2, 5, 1, 3, 2, 5, 2, 5, 1, 1, 1, 5, 3, 1, 3, 1, 5, 3, 3, 4, 1, 1, 4, 4, 1, 3, 3, 2, 5, 5, 1, 1, 1, 1, 3, 1, 5, 2, 1, 3, 5, 1, 4, 3, 1, 3, 1, 1, 3, 1, 1, 1, 1, 1, 1, 5, 1, 1, 5, 5, 2, 1, 5, 1, 4, 1, 1, 5, 1, 1, 1, 5, 5, 5, 1, 4, 5, 1, 3, 1, 2, 5, 1, 1, 1, 5, 1, 1, 4, 1, 1, 2, 3, 1, 3, 4, 1, 2, 1, 4, 3, 1, 2, 4, 1, 5, 1, 1, 1, 1, 1, 3, 4, 1, 1, 5, 1, 1, 3, 1, 1, 2, 1, 3, 1, 2, 1, 1, 3, 3, 4, 5, 3, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 4, 1, 5, 1, 3, 1, 1, 2, 5, 1, 1, 4, 1, 1, 4, 4, 3, 1, 2, 1, 2, 4, 4, 4, 1, 2, 1, 3, 2, 4, 4, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 4, 1, 5, 4, 1, 5, 4, 1, 1, 2, 5, 5, 1, 1, 1, 5}
	fish := []int{3, 4, 3, 1, 2}
	days := 0
	// track number of fish at day 1
	// number of fish at day 2
	// day 3
	// etc
	//
	counts := getCounts(fish)
	fmt.Println("counts", counts)
	fmt.Println("len(counts)", len(counts))
	for days < 18 {
		// shift all counts left
		counts[6] += counts[0]
		counts[len(counts)-1] = counts[0]
		counts[0] = counts[1]
		for i := 2; i < len(counts); i++ {
			counts[i-1] = counts[i]
		}
		fmt.Println("counts", counts)
		days++
	}
	sum := 0
	for i := 0; i < len(counts); i++ {
		sum += counts[i]
	}
	fmt.Println("sum", sum)
}
