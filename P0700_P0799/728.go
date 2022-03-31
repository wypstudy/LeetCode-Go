package main

import (
	"fmt"
	"strconv"
)

func selfDividingNumbers(left int, right int) []int {
	var ans []int
	for i := left; i <= right; i++ {
		check := true
		str := strconv.Itoa(i)
		for _, c := range str {
			div := c - '0'
			if div == 0 {
				check = false
				break
			}
			if i%int(div) != 0 {
				check = false
				break
			}
		}
		if check {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	fmt.Printf("%v", selfDividingNumbers(1, 22))
	fmt.Printf("%v", selfDividingNumbers(47, 85))
}
