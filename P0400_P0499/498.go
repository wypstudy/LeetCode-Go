package main

import "fmt"

func findDiagonalOrder(mat [][]int) []int {
	traveler := Traveler{
		mat:  mat,
		step: 1, max: len(mat) * len(mat[0]),
		x: 0, y: 0, positive: 1,
	}
	ans := make([]int, traveler.max)
	ans[0] = traveler.value()
	for i := 1; i < traveler.max; i++ {
		ans[i] = traveler.next()
	}
	return ans
}

type Traveler struct {
	mat      [][]int
	step     int // step count
	max      int // max step count
	x        int
	y        int
	positive int // 1 mean upper right, -1 mean lower left
}

func (that *Traveler) next() int {
	that.move()
	if that.out() {
		that.turn()
	}
	that.step += 1
	return that.value()
}

func (that *Traveler) move() {
	that.x += -that.positive
	that.y += that.positive
}

func (that *Traveler) turn() {
	that.x += 1
	that.positive = -that.positive
	for that.out() {
		that.move()
	}
}

func (that *Traveler) out() bool {
	return that.x < 0 || that.x >= len(that.mat) || that.y < 0 || that.y >= len(that.mat[0])
}

func (that *Traveler) value() int {
	return that.mat[that.x][that.y]
}

func main() {
	fmt.Printf("%v", findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
