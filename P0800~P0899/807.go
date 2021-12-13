package P0800_P0899

func maxIncreaseKeepingSkyline(grid [][]int) int {
	var x, y [51]int
	var ans = 0
	for i, u := range grid {
		for j, v := range u {
			if v > x[i] {
				x[i] = v
			}
			if v > y[j] {
				y[j] = v
			}
		}
	}
	for i, u := range grid {
		for j, v := range u {
			ans += min(x[i]-v, y[j]-v)
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
