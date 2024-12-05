package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"drjosh.dev/exp"
	"drjosh.dev/exp/emu"
)

// Advent of Code 2016
// Day 25

// Phew! No toggles.

func crim(x string) string {
	if x >= "a" && x <= "z" {
		return fmt.Sprintf("r[%d]", x[0]-'a')
	}
	return x
}

var translators = map[string]emu.TranslatorFunc{
	"cpy": func(_ int, args []string) (string, []int, error) {
		return fmt.Sprintf("%s = %s", crim(args[1]), crim(args[0])), nil, nil
	},
	"inc": func(_ int, args []string) (string, []int, error) {
		return fmt.Sprintf("%s++", crim(args[0])), nil, nil
	},
	"dec": func(_ int, args []string) (string, []int, error) {
		return fmt.Sprintf("%s--", crim(args[0])), nil, nil
	},
	"jnz": func(line int, args []string) (string, []int, error) {
		t := exp.Must(strconv.Atoi(args[1]))
		return fmt.Sprintf("if %s != 0 { goto l%d }", crim(args[0]), line+t), []int{line + t}, nil
	},
	"out": func(_ int, args []string) (string, []int, error) {
		return fmt.Sprintf("if err := send(%s); err != nil { return err }", crim(args[0])), nil, nil
	},
}

func main() {
	program := exp.MustReadLines("inputs/25.txt")
	p := exp.Must(emu.Transpile(program, translators))

	abort := errors.New("abort")

	for a := 1; ; a++ {
		r := []int{a, 0, 0, 0}
		s := false
		c := 0
		p(r, nil, func(x int) error {
			if (x == 1) != s {
				return abort
			}
			s = !s
			c++
			if c > 1_000_000 {
				fmt.Println(a)
				os.Exit(0)
			}
			return nil
		}, nil)
	}
}
