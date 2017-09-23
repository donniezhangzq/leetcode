package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var m map[byte]int = make(map[byte]int)
	var MaxLength int = 0
	var t int = 0
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			//clear map
			for k := t; k < m[s[i]]; k++ {
				delete(m, s[k])
			}
			t = m[s[i]] + 1
			m[s[i]] = i
		} else {
			m[s[i]] = i
			if i-t+1 > MaxLength {
				MaxLength = i - t + 1
			}
		}
	}
	return MaxLength
}

func main() {
	var s1 string = "zwnigfunjwz"
	var r int
	r = lengthOfLongestSubstring(s1)
	fmt.Println(r)
}
