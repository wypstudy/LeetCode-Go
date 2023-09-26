package main

import "fmt"

func passThePillow(n int, time int) int {
	cnt := n - 1
	x := time / cnt
	y := time % cnt
	if x%2 == 0 {
		return 1 + y
	} else {
		return n - y
	}
}

func main() {
	fmt.Println(passThePillow(4, 5))
	fmt.Println(passThePillow(3, 2))
}
