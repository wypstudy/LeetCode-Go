package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func dfs(n *Node, level int, ans *[][]int) {
	if n == nil {
		return
	}
	if level >= len(*ans) {
		*ans = append(*ans, []int{n.Val})
	} else {
		(*ans)[level] = append((*ans)[level], n.Val)
	}
	for _, c := range n.Children {
		dfs(c, level+1, ans)
	}
}

func levelOrder(root *Node) [][]int {
	var solve [][]int
	dfs(root, 0, &solve)
	return solve
}

func main() {
	rt := Node{Val: 1}
	rt.Children = []*Node{{Val: 3}, {Val: 2}, {Val: 4}}
	fmt.Printf("%v", levelOrder(&rt))
}
