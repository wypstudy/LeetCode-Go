package main

import "fmt"

func categorizeBox(length int, width int, height int, mass int) string {
	re := [][]string{
		{"Neither", "Bulky"},
		{"Heavy", "Both"},
	}
	bulky, heavy := 0, 0
	if length >= 10000 || width >= 10000 || height >= 10000 || length*width*height >= 1000000000 {
		bulky = 1
	}
	if mass >= 100 {
		heavy = 1
	}
	return re[heavy][bulky]
}

func main() {
	fmt.Println(categorizeBox(1000, 35, 700, 300))
	fmt.Println(categorizeBox(200, 50, 800, 50))
}
