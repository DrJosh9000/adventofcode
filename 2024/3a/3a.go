package main

import (
	_ "embed"
	"fmt"
	"regexp"

	"drjosh.dev/exp"
)

//go:embed inputs/3.txt
var input string

func main() {
	mulRE := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	sum := 0
	for _, sm := range mulRE.FindAllStringSubmatch(input, -1) {
		sum += exp.MustAtoi(sm[1]) * exp.MustAtoi(sm[2])
	}
	fmt.Println(sum)
}
