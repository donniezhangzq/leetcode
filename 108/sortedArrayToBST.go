package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var s *TreeNode
	fmt.Println(nums[:len(nums)/2])
	fmt.Println(nums[len(nums)/2+1:])
	s = &TreeNode{nums[len(nums)/2], nil, nil}
	fmt.Println(s)
	(*s).Left = sortedArrayToBST(nums[:len(nums)/2])
	(*s).Right = sortedArrayToBST(nums[len(nums)/2+1:])
	return s
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	r := sortedArrayToBST(s)
	fmt.Println(r)
}
