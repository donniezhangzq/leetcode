package main

import "fmt"

func reverse(x int) int {
	const MAX_INT int = int(^uint32(0) >> 1)
	var c, y, t, r, flag int
	r = 0

	if x < 0 {
		flag = 1
		t = 0 - x
	} else {
		flag = 0
		t = x
	}

	c = t / 10
	y = t % 10

	for c+y != 0 {
		if uint(r) > (uint(MAX_INT)-uint(y))/10 {
			return 0
		}

		r = 10*r + y

		t = c
		c = t / 10
		y = t % 10
	}

	if flag == 0 {
		return r
	} else {
		return 0 - r
	}
}

func main() {
	fmt.Println("vim-go")
	r := reverse(-429496729)
	fmt.Println("r:", r)
}
