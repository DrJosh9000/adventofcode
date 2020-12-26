package main

import (
	"fmt"
	"io"
	"os"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m+"\n", p...)
	os.Exit(1)
}

func main() {
	f, err := os.Open("input.8")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	type ins struct {
		opcode string
		operand int
		seen bool
	}
	var prog []*ins
	for {
		in := &ins{}
		_, err := fmt.Fscanf(f, "%s %d\n", &in.opcode, &in.operand)
		if err == io.EOF {
			break
		}		
		if err != nil {
			die("Couldn't fscanf: %v", err)
		}
		prog = append(prog, in)
	}
	pc, acc := 0, 0
	for {
		in := prog[pc]
		if in.seen {
			fmt.Println(acc)
			return
		}
		in.seen = true
		switch in.opcode {
		case "nop":
			pc++
		case "acc":
			acc += in.operand
			pc++
		case "jmp":
			pc += in.operand
		} 
	}
}
