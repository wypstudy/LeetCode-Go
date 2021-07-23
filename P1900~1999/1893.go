package P1900_1999

func isCovered(ranges [][]int, left int, right int) bool {
	// 存储覆盖情况
	var visit [51]bool
	for i := range visit {
		visit[i] = false
	}
	// 覆盖range
	for _, x := range ranges {
		for j := x[0]; j <= x[1]; j++ {
			visit[j] = true
		}
	}
	// 看看是否覆盖
	for i := left; i <= right; i++ {
		if !visit[i] {
			return false
		}
	}
	return true
}
