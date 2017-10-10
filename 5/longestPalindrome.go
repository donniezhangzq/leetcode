package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return s
	}
	var IndexOne int = 0
	var IndexTwo float64 = 0.5
	var LongestPalStr string = string(s[0])
	var r string
	//center one
	for IndexOne < len(s) {
		r = SubStringMaxPalindrome(s, IndexOne, IndexOne)
		if len(r) > len(LongestPalStr) {
			LongestPalStr = r
		}
		IndexOne += 1
	}

	//center two
	for IndexTwo < float64(len(s)-1) {
		if s[int(IndexTwo+0.5)] == s[int(IndexTwo-0.5)] {
			if len(LongestPalStr) < 2 {
				LongestPalStr = s[int(IndexTwo-0.5) : int(IndexTwo+0.5)+1]
			}
			r = SubStringMaxPalindrome(s, int(IndexTwo-0.5), int(IndexTwo+0.5))
			if len(r) > len(LongestPalStr) {
				LongestPalStr = r
			}
		}
		IndexTwo += 1
	}
	return LongestPalStr
}

func SubStringMaxPalindrome(s string, center1 int, center2 int) string {
	var r string = string(s[center1 : center2+1])
	fmt.Println("center1:", center1, "center2", center2)
	for i := 1; center1-i >= 0 && center2+i < len(s); i++ {
		if s[center1-i] == s[center2+i] {
			r = s[center1-i : center2+i+1]
		} else {
			return r
		}
	}
	return r
}

func main() {
	var s string = "abcda"
	r := longestPalindrome(s)
	fmt.Println("result is:", r)
}
