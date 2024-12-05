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
	args []any
}

func main() {
	var prog []inst
	exp.MustForEachLineIn("inputs/18.txt", func(line string) {
		tokens := strings.Fields(line)
		if len(tokens) < 2 {
			log.Fatalf("Too few tokens in line %q [%d < 2]", line, len(tokens))
		}
		var in inst
		in.opcode = tokens[0]
		for _, arg := range tokens[1:] {
			n, err := strconv.Atoi(arg)
			if err != nil {
				in.args = append(in.args, arg)
				continue
			}
			in.args = append(in.args, n)
		}
		prog = append(prog, in)
	})

	regs := make(map[string]int)
	eval := func(arg any) int {
		switch x := arg.(type) {
		case string: return regs[x]
		case int: return x
		}
		log.Fatalf("Bad arg type %T [must be string or int]", arg)
		return -666
	}
	lastSnd := -666
	ip := 0
	for {
		in := prog[ip]
		switch in.opcode {
		case "snd":
			lastSnd = eval(in.args[0])
		case "set":
			regs[in.args[0].(string)] = eval(in.args[1])
		case "add":
			regs[in.args[0].(string)] += eval(in.args[1])
		case "mul":
			regs[in.args[0].(string)] *= eval(in.args[1])
		case "mod":
			regs[in.args[0].(string)] %= eval(in.args[1])
		case "rcv":
			if eval(in.args[0]) != 0 {
				fmt.Println(lastSnd)
				return
			}
		case "jgz":
			if eval(in.args[0]) > 0 {
				ip += eval(in.args[1])
				continue
			}
		}
		ip++
	}
}
