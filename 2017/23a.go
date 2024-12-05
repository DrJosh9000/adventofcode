package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"drjosh.dev/exp"
)

type inst struct {
	opcode string
	arg    [2]any
}

func main() {
	var prog []inst
	exp.MustForEachLineIn("inputs/23.txt", func(line string) {
		tokens := strings.Fields(line)
		if len(tokens) < 2 {
			log.Fatalf("Too few tokens in line %q [%d < 2]", line, len(tokens))
		}
		var in inst
		in.opcode = tokens[0]
		for i, a := range tokens[1:] {
			n, err := strconv.Atoi(a)
			if err != nil {
				in.arg[i] = a[0] - 'a'
				continue
			}
			in.arg[i] = n
		}
		prog = append(prog, in)
	})

	regs := make([]int, 8)
	eval := func(arg any) int {
		switch x := arg.(type) {
		case byte:
			return regs[x]
		case int:
			return x
		}
		log.Fatalf("Bad arg type %T [must be byte or int]", arg)
		return -666
	}
	mulCount := 0
	for ip := 0; ip >= 0 && ip < len(prog); {
		in := prog[ip]
		switch in.opcode {
		case "set":
			regs[in.arg[0].(byte)] = eval(in.arg[1])
		case "sub":
			regs[in.arg[0].(byte)] -= eval(in.arg[1])
		case "mul":
			regs[in.arg[0].(byte)] *= eval(in.arg[1])
			mulCount++
		case "jnz":
			if eval(in.arg[0]) != 0 {
				ip += eval(in.arg[1])
				continue
			}
		}
		ip++
	}
	fmt.Println(mulCount)
	fmt.Println(regs)
}
