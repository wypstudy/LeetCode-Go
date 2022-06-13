package main

import "sort"

func heightChecker(heights []int) int {
	expects := make([]int, len(heights))
	copy(expects, heights)
	sort.Ints(expects)
	count := 0
	for i, e := range expects {
		if heights[i] != e {
			count += 1
		}
	}
	return count
}
