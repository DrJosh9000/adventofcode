package main

import (
	"fmt"

	"drjosh.dev/exp"
)

func main() {
	const mod = 2147483647
	const mask = 0xffff
	factor := [2]int{16807, 48271}
	div := [2]int{4, 8}
	gch := [2]chan int{make(chan int), make(chan int)}
	gen := func(x, f, d int, ch chan int) {
		for {
			x = (x * f) % mod
			if x%d == 0 {
				ch <- x
			}
		}
	}
	exp.MustForEachLineIn("inputs/15.txt", func(line string) {
		var s int
		var g rune
		exp.Must(fmt.Sscanf(line, "Generator %c starts with %d", &g, &s))
		n := g - 'A'
		go gen(s, factor[n], div[n], gch[n])
	})

	judge := 0
	for i := 0; i < 5_000_000; i++ {
		a, b := <-gch[0], <-gch[1]
		if a&mask == b&mask {
			judge++
		}
	}
	fmt.Println(judge)
}
