package main

import (
	"fmt"
	"strings"
)

func lengthLongestPath(input string) int {
	inputs := strings.Split(input, "\n")
	return find(0, 0, len(inputs), "", inputs)
}

func find(tc int, si int, sj int, path string, ss []string) int {
	ans := 0
	split := ""
	if tc > 0 {
		split = "/"
	}
	for i := si; i < sj; i++ {
		if strings.Count(ss[i], "\t") == tc {
			// same level
			if !strings.Contains(ss[i], ".") {
				// dir action
				j := i + 1
				for j < sj && strings.Count(ss[j], "\t") > tc {
					j++
				}
				newAns := find(tc+1, i, j, path+split+strings.Trim(ss[i], "\t*"), ss)
				if newAns > ans {
					ans = newAns
				}
			} else {
				// file action
				// fmt.Println(path + split + strings.Trim(ss[i], "\t*"))
				newAns := len(path + split + strings.Trim(ss[i], "\t*"))
				if newAns > ans {
					ans = newAns
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Printf("%v\n", lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"))
	fmt.Printf("%v\n", lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"))
	fmt.Printf("%v\n", lengthLongestPath("a"))
	fmt.Printf("%v\n", lengthLongestPath("rzzmf\nv\n\tix\n\t\tiklav\n\t\t\ttqse\n\t\t\t\ttppzf\n\t\t\t\t\tzav\n\t\t\t\t\t\tkktei\n\t\t\t\t\t\t\thhmav\n\t\t\t\t\t\t\t\tbzvwf.txt"))
}
