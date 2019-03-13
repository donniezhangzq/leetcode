package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func myAtoi(str string) int {
	const (
		maxInt = ^(1 << 31)
		minInt = 1 - maxInt
		zero = 0
	)

	chars := []byte(str)
	if len(chars) < 1 {
		return 0
	}

}