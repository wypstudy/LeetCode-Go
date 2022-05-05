package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	ans, product, i := 0, 1, 0
	for j, num := range nums {
		product *= num
		for i <= j && product >= k {
			product /= nums[i]
			i++
		}
		//fmt.Printf("i=%d,j=%d,product=%d\n", i, j, product)
		ans += j - i + 1
	}
	return ans
}

func main() {
	fmt.Printf("%v\n", numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}
