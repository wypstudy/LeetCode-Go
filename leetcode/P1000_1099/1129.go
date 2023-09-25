package main

type Status struct {
	pos       int8
	step      int8
	nextColor int8 // 0 = both, 1 = red, 2 = blue
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	visited := make([]bool, n)
	answer := make([]int, n)
	redMap := make([][]bool, n)
	blueMap := make([][]bool, n)
	for i := 0; i < n; i++ {
		answer[i] = -1
		redMap[i] = make([]bool, n)
		blueMap[i] = make([]bool, n)
	}
	for i := 0; i < len(redEdges); i++ {
		redMap[redEdges[i][0]][redEdges[i][1]] = true
	}
	for i := 0; i < len(blueEdges); i++ {
		blueMap[blueEdges[i][0]][blueEdges[i][1]] = true
	}
	statusQueue := make([]Status, 0)
	visited[0] = true
	answer[0] = 0
	statusQueue = append(statusQueue, Status{0, 0, 0})
	statusQueueIndex := 0
	for statusQueueIndex < len(statusQueue) {

	}

	return answer
}
