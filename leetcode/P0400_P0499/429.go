package main

type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

func dfs(n *NTreeNode, level int, ans *[][]int) {
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

func levelOrder(root *NTreeNode) [][]int {
	var solve [][]int
	dfs(root, 0, &solve)
	return solve
}

// func main() {
// 	rt := NTreeNode{Val: 1}
// 	rt.Children = []*NTreeNode{{Val: 3}, {Val: 2}, {Val: 4}}
// 	fmt.Printf("%v", levelOrder(&rt))
// }
