package main

import "fmt"

func numTrees(n int) int {
	if n <= 0 {
		return 0
	}
	var res []int = make([]int, n+1)

	res[0] = 1
	res[1] = 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			res[i] += res[j] * res[i-j-1]
		}
	}
	return res[n]
}

func main() {
	r := numTrees(4)
	fmt.Println(r)
}
