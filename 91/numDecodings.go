package main

import "fmt"
import "strconv"

func numDecodings(s string) int {
	if s == "" {
		return 0
	}
	rs := []rune(s)
	if rs[0] == '0' {
		return 0
	} else if len(rs) == 1 {
		return 1
	}
	var dp []int = make([]int, len(rs))
	dp[0] = 1
	if it, _ := strconv.Atoi(string(rs[0:2])); it >= 11 && it <= 26 {
		if rs[1] == '0' {
			dp[1] = 1
		} else {
			dp[1] = 2
		}
	} else if rs[1] == '0' {
		if it > 26 {
			return 0
		} else {
			dp[1] = 1
		}
	} else {
		dp[1] = 1
	}
	for i := 2; i < len(rs); i++ {
		i1, err1 := strconv.Atoi(string(rs[i]))
		i2, err2 := strconv.Atoi(string(rs[i-1 : i+1]))
		fmt.Println(i1, i2, "\n")
		if err1 != nil || err2 != nil {
			return 0
		}
		if i1 >= 1 && i1 <= 9 {
			dp[i] += dp[i-1]
			fmt.Println("dp[i]", dp[i], "dp[i-1]", dp[i-1])
		}
		if i2 >= 10 && i2 <= 26 {
			dp[i] += dp[i-2]
			fmt.Println("dp[i]", dp[i], "dp[i-2]", dp[i-2])
		}
		fmt.Printf("dp[%d] is %d\n", i, dp[i])
	}
	return dp[len(rs)-1]
}

func main() {
	a := "110"
	fmt.Printf("%s 's result is %d\n", a, numDecodings(a))
}
