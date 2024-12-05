package main

import (
	"fmt"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

func main() {
	fmt.Println(algo.Sum(exp.MustReadInts("inputs/1.txt", "\n")))
}
