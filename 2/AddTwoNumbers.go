package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var index int = 0
	var sr []int = make([]int, 2)
	for {
		if l1 == nil && l2 == nil {
			break
		} else if l1 == nil {
			sr[index] += (*l2).Val
		} else if l2 == nil {
			sr[index] += (*l1).Val
		} else {
			if len(sr) < index+1 {
				sr = append(sr, (*l1).Val+(*l2).Val+sr[index])
			} else {
				sr[index] += (*l1).Val + (*l2).Val
			}
		}

		if sr[index] >= 10 {
			Add := sr[index] / 10
			sr[index] = sr[index] % 10
			sr = append(sr, Add)
		}
		index++

		l1 = (*l1).Next
		l2 = (*l2).Next
	}
	var lr *ListNode
	var lt *ListNode
	lt = &ListNode{sr[0], nil}
	lr = lt
	i := 1
	for i < len(sr) {
		(*lt).Next = &ListNode{sr[i], nil}
		lt = (*lt).Next
		i++
	}
	return lr
}

func main() {
	l1 := &ListNode{9, &ListNode{8, &ListNode{7, nil}}}
	l2 := &ListNode{6, &ListNode{7, &ListNode{4, nil}}}
	lt := addTwoNumbers(l1, l2)
	fmt.Println(lt)
}
