package main

func decodeMessage(key string, message string) string {
	charMap := make(map[rune]rune, 0)
	for _, char := range key {
		if char == ' ' {
			continue
		}
		if _, ok := charMap[char]; !ok {
			val := rune('a' + len(charMap))
			charMap[char] = val
		}
	}
	res := ""
	for _, char := range message {
		if val, ok := charMap[char]; ok {
			res += string(val)
		} else {
			res += string(char)
		}
	}
	return res
}
