package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var sr []int
	var Add int = 0
	for {
		if l1 == nil && l2 == nil {
			if Add != 0 {
				sr = append(sr, Add)
			}
			break
		} else if l1 == nil {
			sr = append(sr, (*l2).Val)
		} else if l2 == nil {
			sr = append(sr, (*l1).Val)
		} else {
			sr = append(sr, (*l1).Val+(*l2).Val)
		}

		tAdd := (sr[len(sr)-1] + Add) / 10
		sr[len(sr)-1] = (sr[len(sr)-1] + Add) % 10
		Add = tAdd

		if l1 != nil {
			l1 = (*l1).Next
		}
		if l2 != nil {
			l2 = (*l2).Next
		}
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
	l1 := &ListNode{9, &ListNode{9, nil}}
	l2 := &ListNode{1, nil}
	lt := addTwoNumbers(l1, l2)
	fmt.Println(lt)
}
