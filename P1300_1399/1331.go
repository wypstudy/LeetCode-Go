package main

import "sort"

func arrayRankTransform(arr []int) []int {
	// make set
	set := make(map[int]bool)
	for _, num := range arr {
		set[num] = true
	}

	// set to array
	afterSet := make([]int, 0)
	for k, _ := range set {
		afterSet = append(afterSet, k)
	}

	// sort the array
	sort.Ints(afterSet)

	// gen index
	indexMap := make(map[int]int)
	for index, num := range afterSet {
		indexMap[num] = index + 1
	}

	// index replace
	re := make([]int, len(arr))
	for index, num := range arr {
		re[index] = indexMap[num]
	}

	return re
}

func main() {

}
