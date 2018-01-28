package main

import "fmt"

func maxSubArray(nums []int) int {
	var s, st, e, r, rt, t_sum1 int = 0, 0, 0, 0, 0, 0
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	r = nums[0]

	for i := 1; i < len(nums); i++ {
		t_sum1 = 0
		st = s
		rt = r
		for j := i; j > e; j-- {
			t_sum1 += nums[j]
			if t_sum1 >= r {
				s = j
				r = t_sum1
			}
			fmt.Println("s,e,r", s, e, r)
			fmt.Println("j,nums[j],t_sum1", j, nums[j], t_sum1)
		}
		if rt+t_sum1 >= r {
			r = rt + t_sum1
			e = i
			s = st
		} else if s > e {
			e = i
		}
		fmt.Println("s,e,r", s, e, r)
		fmt.Println("i,nums[i],t_sum1", i, nums[i], t_sum1)
	}
	fmt.Println(nums[s : e+1])
	return r
}

func main() {
	var s []int = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	r := maxSubArray(s)
	fmt.Println(r)
}
