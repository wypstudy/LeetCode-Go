package main

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minFlipsMonoIncr(s string) int {
	dp0, dp1 := 0, 0
	for _, c := range s {
		ndp0, ndp1 := dp0, min(dp0, dp1)
		if c == '0' {
			ndp1 += 1
		} else {
			ndp0 += 1
		}
		dp0, dp1 = ndp0, ndp1
	}
	return min(dp0, dp1)
}
