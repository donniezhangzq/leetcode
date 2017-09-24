package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result []int = make([]int, len(nums1)+len(nums2))
	var IndexNums1 int = 0
	var IndexNums2 int = 0

	for i := 0; i < len(nums1)+len(nums2); i++ {
		if IndexNums1 >= len(nums1) && IndexNums2 >= len(nums2) {
			break
		} else if IndexNums1 >= len(nums1) {
			result[i] = nums2[IndexNums2]
			IndexNums2++
		} else if IndexNums2 >= len(nums2) {
			result[i] = nums1[IndexNums1]
			IndexNums1++
		} else {
			if nums1[IndexNums1] < nums2[IndexNums2] {
				result[i] = nums1[IndexNums1]
				IndexNums1++
			} else {
				result[i] = nums2[IndexNums2]
				IndexNums2++
			}
		}
	}
	if len(result)%2 == 0 {
		return float64(result[len(result)/2-1]+result[len(result)/2]) / 2
	} else {
		return float64(result[len(result)/2])
	}

}

func main() {
	var nums1 []int = []int{1, 2}
	var nums2 []int = []int{3, 4}
	r := findMedianSortedArrays(nums1, nums2)
	fmt.Println(r)
}
