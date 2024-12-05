package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2016
// Day 23, part a

func main() {
	type inst struct {
		opcode string
		args   []any
	}
	var program []inst

	for _, line := range exp.MustReadLines("inputs/23.txt") {
		var i inst
		t := strings.Fields(line)
		i.opcode = t[0]
		for _, a := range t[1:] {
			n, err := strconv.Atoi(a)
			if err != nil {
				i.args = append(i.args, a[0])
				continue
			}
			i.args = append(i.args, n)
		}
		program = append(program, i)
	}

	toggle := map[string]string{
		"cpy": "jnz",
		"jnz": "cpy",
		"inc": "dec",
		"dec": "inc",
		"tgl": "inc",
	}

	reg := []int{exp.Must(strconv.Atoi(os.Args[1])), 0, 0, 0}

	eval := func(a any) int {
		switch x := a.(type) {
		case int:
			return x
		case byte:
			return reg[x-'a']
		}
		panic(fmt.Sprintf("unsupported arg type %T", a))
	}

	for ip := 0; ip < len(program); {
		i := program[ip]
		//fmt.Println(i.opcode, i.args)
		switch i.opcode {
		case "cpy":
			r, ok := i.args[1].(byte)
			if !ok {
				ip++
				continue
			}
			reg[r-'a'] = eval(i.args[0])
			ip++

		case "jnz":
			j := 1
			if eval(i.args[0]) != 0 {
				j = eval(i.args[1])
			}
			ip += j

		case "inc":
			r, ok := i.args[0].(byte)
			if !ok {
				ip++
				continue
			}
			reg[r-'a']++
			ip++

		case "dec":
			r, ok := i.args[0].(byte)
			if !ok {
				ip++
				continue
			}
			reg[r-'a']--
			ip++

		case "tgl":
			j := ip + eval(i.args[0])
			if j < 0 || j >= len(program) {
				ip++
				continue
			}
			program[j].opcode = toggle[program[j].opcode]
			ip++
		}
	}

	fmt.Println(reg[0])
}
