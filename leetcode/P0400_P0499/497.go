package main

import (
	"math/rand"
	"sort"
)

type Solution struct {
	rects   [][]int
	sizeSum []int
	count   int
}

func Constructor(rects [][]int) Solution {
	count := len(rects)
	sizeSum := make([]int, count+1)
	for i, r := range rects {
		a, b, x, y := r[0], r[1], r[2], r[3]
		size := (x - a + 1) * (y - b + 1)
		sizeSum[i+1] = sizeSum[i] + size
	}
	return Solution{rects: rects, sizeSum: sizeSum, count: count}
}

func (this *Solution) Pick() []int {
	chooseSize := rand.Intn(this.sizeSum[this.count])
	chooseIndex := sort.SearchInts(this.sizeSum, chooseSize+1) - 1
	realSize := chooseSize - this.sizeSum[chooseIndex]
	r := this.rects[chooseIndex]
	a, b, y := r[0], r[1], r[3]
	u := realSize/(y-b+1) + a
	v := realSize%(y-b+1) + b
	return []int{u, v}
}
