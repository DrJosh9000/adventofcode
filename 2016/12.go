package main

import (
	"fmt"
	"strconv"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/emu"
)

// Advent of Code 2016
// Day 12

func crim(x string) string {
	if x >= "a" && x <= "z" {
		return fmt.Sprintf("r[%d]", x[0]-'a')
	}
	return x
}

var translators = map[string]emu.TranslatorFunc{
	"cpy": func(_ int, args ...string) (string, []int, error) {
		return fmt.Sprintf("%s = %s", crim(args[1]), crim(args[0])), nil, nil
	},
	"inc": func(_ int, args ...string) (string, []int, error) {
		return fmt.Sprintf("%s++", crim(args[0])), nil, nil
	},
	"dec": func(_ int, args ...string) (string, []int, error) {
		return fmt.Sprintf("%s--", crim(args[0])), nil, nil
	},
	"jnz": func(line int, args ...string) (string, []int, error) {
		t := exp.Must(strconv.Atoi(args[1]))
		return fmt.Sprintf("if %s != 0 { goto l%d }", crim(args[0]), line+t), []int{line + t}, nil
	},
}

func main() {
	program := exp.MustReadLines("inputs/12.txt")
	p := exp.Must(emu.Transpile(program, translators))
	r := []int{0, 0, 0, 0}
	p(r, nil, nil, nil)
	fmt.Println(r[0])

	r = []int{0, 0, 1, 0}
	p(r, nil, nil, nil)
	fmt.Println(r[0])
}
