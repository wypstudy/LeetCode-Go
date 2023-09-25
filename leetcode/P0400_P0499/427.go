package main

import "fmt"

/**
 * Definition for a QuadTree node.
 **/

type QuadTreeNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadTreeNode
	TopRight    *QuadTreeNode
	BottomLeft  *QuadTreeNode
	BottomRight *QuadTreeNode
}

func construct(grid [][]int) *QuadTreeNode {
	// check if the grid is a leaf
	x := grid[0][0]
	l := len(grid)
	m := l / 2
	isLeafCheck := true
	for i := 0; isLeafCheck && i < l; i++ {
		for j := 0; isLeafCheck && j < l; j++ {
			if grid[i][j] != x {
				isLeafCheck = false
			}
		}
	}
	node := QuadTreeNode{
		IsLeaf:      isLeafCheck,
		Val:         x == 1,
		TopLeft:     nil,
		TopRight:    nil,
		BottomLeft:  nil,
		BottomRight: nil,
	}
	// gen sub-node when isLeaf = false
	if !isLeafCheck {
		node.TopLeft = construct(slice2d(grid, 0, m, 0, m))
		node.TopRight = construct(slice2d(grid, 0, m, m, l))
		node.BottomLeft = construct(slice2d(grid, m, l, 0, m))
		node.BottomRight = construct(slice2d(grid, m, l, m, l))
	}
	return &node
}

func slice2d(grid [][]int, rowStart int, rowEnd int, colStart int, colEnd int) [][]int {
	newGrid := make([][]int, rowEnd-rowStart)
	copy(newGrid, grid[rowStart:rowEnd])
	for index, row := range newGrid {
		newGrid[index] = row[colStart:colEnd]
	}
	return newGrid
}

func main() {
	grid := [][]int{{1, 1, 0, 0}, {0, 0, 1, 1}, {1, 1, 0, 0}, {0, 0, 1, 1}}
	fmt.Printf("%v\n", construct(grid))
}
