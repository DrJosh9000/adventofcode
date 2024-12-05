package main

import (
	"fmt"
	"log"
	"strings"

	"drjosh.dev/exp"
)

func main() {
	type input struct {
		state  byte
		curVal int
	}
	type output struct {
		nextState byte
		writeVal  int
		nextSlot  int
	}
	program := make(map[input]output)
	var (
		state byte
		steps int
		in    input
		out   output
	)
	exp.MustForEachLineIn("inputs/25.txt", func(line string) {
		//fmt.Println(line)
		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "Begin in state"):
			exp.Must(fmt.Sscanf(line, "Begin in state %c.", &state))
		case strings.HasPrefix(line, "Perform a diagnostic checksum after"):
			exp.Must(fmt.Sscanf(line, "Perform a diagnostic checksum after %d steps.", &steps))
		case strings.HasPrefix(line, "In state"):
			exp.Must(fmt.Sscanf(line, "In state %c:", &in.state))
		case strings.HasPrefix(line, "If the current value is"):
			exp.Must(fmt.Sscanf(line, "If the current value is %d:", &in.curVal))
		case strings.HasPrefix(line, "- Write the value"):
			exp.Must(fmt.Sscanf(line, "- Write the value %d.", &out.writeVal))
			program[in] = out
		case strings.HasPrefix(line, "- Move one slot to the"):
			var s string
			exp.Must(fmt.Sscanf(line, "- Move one slot to the %s", &s))
			out.nextSlot = -1
			if s == "right." {
				out.nextSlot = 1
			}
			program[in] = out
		case strings.HasPrefix(line, "- Continue with state"):
			exp.Must(fmt.Sscanf(line, "- Continue with state %c.", &out.nextState))
			program[in] = out
		}
	})

	tape := make(map[int]int)
	head := 0
	for i := 0; i < steps; i++ {
		out, ok := program[input{state: state, curVal: tape[head]}]
		if !ok {
			log.Fatalf("No matching input on step %d", i)
		}
		tape[head] = out.writeVal
		state = out.nextState
		head += out.nextSlot
	}

	count := 0
	for _, v := range tape {
		count += v
	}

	fmt.Println(count)
}
