package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"

	"drjosh.dev/exp"
)

//go:embed inputs/3.txt
var input string

func main() {
	opRE := regexp.MustCompile(`(do|don't|mul)\(([\d,]*)\)`)
	sum := 0
	enable := true
	for _, sm := range opRE.FindAllStringSubmatch(input, -1) {
		switch sm[1] {
		case "mul":
			if !enable {
				continue
			}
			args := strings.Split(sm[2], ",")
			if len(args) != 2 {
				continue
			}
			sum += exp.MustAtoi(args[0]) * exp.MustAtoi(args[1])

		case "do":
			if sm[2] != "" {
				continue
			}
			enable = true

		case "don't":
			if sm[2] != "" {
				continue
			}
			enable = false
		}
	}
	fmt.Println(sum)
}
