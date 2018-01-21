package main

import "fmt"
import "time"

func climbStairs(c int) int {
	a := 1
	b := 2
	r := 0
	if c == 1 {
		return 1
	} else if c == 2 {
		return 2
	} else {
		for i := 3; i <= c; i++ {
			r = a + b
			a = b
			b = r
		}
		return r
	}
}

func main() {
	start := time.Now()
	fmt.Printf("climbStairs steps is %d\n", climbStairs(42))
	fmt.Printf("climbStairs steps is %d\n", climbStairs(43))
	fmt.Printf("climbStairs steps is %d\n", climbStairs(44))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("process time is %s", delta)
}
