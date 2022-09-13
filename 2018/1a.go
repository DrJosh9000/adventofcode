package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

func main() {
	fmt.Println(algo.Sum(exp.MustReadInts("inputs/1.txt", "\n")))
}
