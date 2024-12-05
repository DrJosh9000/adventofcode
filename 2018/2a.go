package main

import (
	"fmt"

	"drjosh.dev/exp"
)

func main() {
	two, three := 0, 0
	exp.MustForEachLineIn("inputs/2.txt", func(line string) {
		m := make(map[rune]int)
		for _, r := range line {
			m[r]++
		}
		c2, c3 := false, false
		for _, v := range m {
			c2 = c2 || (v == 2)
			c3 = c3 || (v == 3)
		}
		if c2 { two++ }
		if c3 { three++ }
	})
	fmt.Println(two * three)
}