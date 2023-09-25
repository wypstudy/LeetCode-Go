package main

func rotateString(s string, goal string) bool {
	n := len(s)
	for i, _ := range s {
		ns := s[n-i:n] + s[:n-i]
		if ns == goal {
			return true
		}
	}
	return false
}

// func main() {
// 	println(rotateString("abcde", "cdeab"))
// 	println(rotateString("abcde", "abced"))
// }
