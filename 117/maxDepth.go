package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepth := recurTreeDepth(root, 0)
	return maxDepth
}

func recurTreeDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	maxDepthLeft := recurTreeDepth((*root).Left, depth+1)
	maxDepthRight := recurTreeDepth((*root).Right, depth+1)
	if maxDepthLeft > maxDepthRight {
		return maxDepthLeft
	} else {
		return maxDepthRight
	}
}

func main() {
	t := &TreeNode{1, &TreeNode{2, nil, &TreeNode{3, nil, nil}}, nil}
	maxDepth := maxDepth(t)
	fmt.Println("maxDEpth", maxDepth)
}
