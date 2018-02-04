package main

import "fmt"

func maxProduct(nums []int) int {
	var rMax, rMaxt, rMin, rMint, max int

	if len(nums) < 1 {
		return 0
	}
	rMin = nums[0]
	rMax = nums[0]
	rMint = nums[0]
	rMaxt = nums[0]
	max = nums[0]

	for i := 1; i < len(nums); i++ {
		rMax = findMax(rMaxt*nums[i], rMint*nums[i], nums[i])
		rMin = findMin(rMaxt*nums[i], rMint*nums[i], nums[i])
		rMaxt = rMax
		rMint = rMin
		if rMax > max {
			max = rMax
		}
	}
	return max
}

func findMax(a int, b int, c int) int {
	r := a
	if r < a {
		r = a
	}
	if r < b {
		r = b
	}
	if r < c {
		r = c
	}
	return r
}
func findMin(a int, b int, c int) int {
	r := a
	if r > a {
		r = a
	}
	if r > b {
		r = b
	}
	if r > c {
		r = c
	}
	return r
}

func main() {
	var s []int = []int{-4, -3, -2}
	r := maxProduct(s)
	fmt.Println(r)
}
