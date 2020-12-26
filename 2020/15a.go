package main

import "fmt"

func main() {
	mem := map[int]int{
		12: 1,
		1:  2,
		16: 3,
		3:  4,
		11: 5,
		0:  6,
	}
	last := 0
	for turn := 8; turn <= 2020; turn++ {
		prev, old := mem[last]
		mem[last] = turn - 1
		if old {
			last = turn - 1 - prev
		} else {
			last = 0
		}
		fmt.Println(last)
	}
}
