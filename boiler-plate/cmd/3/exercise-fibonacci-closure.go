package main

import "fmt"

func fibonacci() func() int {
	beforeBefore := 0
	before := 0
	return func() int {
		sum := beforeBefore + before
		if sum == 0 {
			before = 1
		} else {
			beforeBefore = before
			before = sum
		}
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
