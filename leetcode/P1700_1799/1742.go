package main

import "fmt"

func countBalls(lowLimit int, highLimit int) int {
	box := make([]int, 50)
	max := 0
	for i := lowLimit; i <= highLimit; i++ {
		boxIndex := ballToBox(i)
		box[boxIndex] += 1
		if box[boxIndex] > max {
			max = box[boxIndex]
		}
	}
	return max
}

func ballToBox(x int) int {
	index := 0
	for {
		index += x % 10
		x /= 10
		if x <= 0 {
			break
		}
	}
	return index
}

func main() {
	fmt.Println(countBalls(1, 10))
}
