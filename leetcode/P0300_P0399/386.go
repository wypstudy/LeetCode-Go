package main

import "fmt"

func gen(x int, n int, ans *[]int) {
	if x <= n {
		*ans = append(*ans, x)
		gen(x*10, n, ans)
		if x%10 < 9 {
			gen(x+1, n, ans)
		}
	}
}

func lexicalOrder(n int) []int {
	var ans []int
	gen(1, n, &ans)
	return ans
}

func main() {
	fmt.Printf("%v", lexicalOrder(13))
}
