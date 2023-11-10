package main

import (
	"fmt"
	"math"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	lenp := len(potions)
	cpPotions := make([]int, lenp)
	copy(cpPotions, potions)
	sort.Ints(cpPotions)
	ans := make([]int, len(spells))
	for idx, spell := range spells {
		minPotion := int(math.Ceil(float64(success) / float64(spell)))
		pos := sort.Search(lenp, func(i int) bool { return cpPotions[i] >= minPotion })
		ans[idx] = lenp - pos
	}
	return ans
}

func main() {
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
	fmt.Println(successfulPairs([]int{3, 1, 2}, []int{8, 5, 8}, 16))
}
