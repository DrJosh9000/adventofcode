package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/DrJosh9000/exp"
)

type inst struct {
	opcode string
	args []any
}

// A goroutine deadlock is fatal, not a panic, so can't be caught in a defer.
// So (atomically!) count the number of goroutines in recv state separately.
// If both are trying to receive, and both of the channels are empty at that
// point, the programs are deadlocking in the intended fashion!

var rcvCount atomic.Int32

func run(prog []inst, id int, inch <-chan int, outch chan<- int) {
	regs := map[string]int{
		"p": id,
	}
	eval := func(arg any) int {
		switch x := arg.(type) {
		case string: return regs[x]
		case int: return x
		}
		log.Fatalf("Bad arg type %T [must be string or int]", arg)
		return -666
	}
	sndCount := 0
	if id == 1 {
		defer func() {
			fmt.Println(sndCount)
		}()
	}
	for ip := 0; ip < len(prog); {
		in := prog[ip]
		switch in.opcode {
		case "snd":
			sndCount++
			outch <- eval(in.args[0])
		case "set":
			regs[in.args[0].(string)] = eval(in.args[1])
		case "add":
			regs[in.args[0].(string)] += eval(in.args[1])
		case "mul":
			regs[in.args[0].(string)] *= eval(in.args[1])
		case "mod":
			regs[in.args[0].(string)] %= eval(in.args[1])
		case "rcv":
			rcvCount.Add(1)
			if rcvCount.Load() >= 2 && len(inch) == 0 && len(outch) == 0 {
				return
			}
			regs[in.args[0].(string)] = <-inch
			rcvCount.Add(-1)
		case "jgz":
			if eval(in.args[0]) > 0 {
				ip += eval(in.args[1])
				continue
			}
		}
		ip++
	}
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
	
	p0in := make(chan int, 10000)
	p1in := make(chan int, 10000)
	go run(prog, 0, p0in, p1in)
	run(prog, 1, p1in, p0in)
}