package main

import (
	"math"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) (ans [][]int) {
	type data struct{ col, row, val int }
	var nodes []data
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}
		nodes = append(nodes, data{col, row, node.Val})
		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}
	dfs(root, 0, 0)
	sort.Slice(nodes, func(i, j int) bool {
		x, y := nodes[i], nodes[j]
		return x.col < y.col ||
			(x.col == y.col && x.row < y.row) ||
			(x.col == y.col && x.row == y.row && x.val < y.val)
	})
	lastCol := math.MinInt32
	for _, node := range nodes {
		if node.col != lastCol {
			lastCol = node.col
			ans = append(ans, nil)
		}
		ans[len(ans)-1] = append(ans[len(ans)-1], node.val)
	}
	return
}
