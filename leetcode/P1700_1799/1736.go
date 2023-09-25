package main

import "fmt"

func maximumTime(time string) string {
	max := []rune(time)
	if max[0] == '?' {
		if max[1] == '?' || max[1] < '4' {
			max[0] = '2'
		} else {
			max[0] = '1'
		}
	}
	if max[1] == '?' {
		if max[0] == '2' {
			max[1] = '3'
		} else {
			max[1] = '9'
		}
	}
	if max[3] == '?' {
		max[3] = '5'
	}
	if max[4] == '?' {
		max[4] = '9'
	}
	return string(max)
}

func main() {
	fmt.Println(maximumTime("2?:2?"))
}
