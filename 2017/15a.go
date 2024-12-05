package main

import (
	"fmt"

	"drjosh.dev/exp"
)

func main() {
	const mod = 2147483647
	const mask = 0xffff
	factor := [2]int{16807, 48271}
	var gen [2]int
	exp.MustForEachLineIn("inputs/15.txt", func(line string) {
		var s int
		var g rune
		exp.Must(fmt.Sscanf(line, "Generator %c starts with %d", &g, &s))
		gen[g-'A'] = s
	})

	judge := 0
	for i := 0; i < 40_000_000; i++ {
		for j := range gen {
			gen[j] = (gen[j] * factor[j]) % mod
		}
		if gen[0]&mask == gen[1]&mask {
			judge++
		}
	}
	fmt.Println(judge)
}
