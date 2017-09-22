package main

import "fmt"

func towSum(nums []int, target int) []int {
    k := 0
    for k < len(nums) {
        for key,value := range nums[k+1:] {
            if (value + nums[k] == target) {
                return []int{k, key+k+1}
            }
        }
        k++
    }
    return []int{} 
}

func main() {
    var nums[]int = []int{3,2,4}
    var target int = 6
    result := towSum(nums, target)
    fmt.Println(result[0], result[1])
}

