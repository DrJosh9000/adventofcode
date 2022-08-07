package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/DrJosh9000/exp"
)

type state [6]int

func (s state) addr(a, b, c int) state {
	s[c] = s[a] + s[b]
	return s
}

func (s state) addi(a, b, c int) state {
	s[c] = s[a] + b
	return s
}

func (s state) mulr(a, b, c int) state {
	s[c] = s[a] * s[b]
	return s
}

func (s state) muli(a, b, c int) state {
	s[c] = s[a] * b
	return s
}

func (s state) banr(a, b, c int) state {
	s[c] = s[a] & s[b]
	return s
}

func (s state) bani(a, b, c int) state {
	s[c] = s[a] & b
	return s
}

func (s state) borr(a, b, c int) state {
	s[c] = s[a] | s[b]
	return s
}

func (s state) bori(a, b, c int) state {
	s[c] = s[a] | b
	return s
}

func (s state) setr(a, _, c int) state {
	s[c] = s[a]
	return s
}

func (s state) seti(a, _, c int) state {
	s[c] = a
	return s
}

func (s state) gtir(a, b, c int) state {
	if a > s[b] {
		s[c] = 1
	} else {
		s[c] = 0
	}
	return s
}

func (s state) gtri(a, b, c int) state {
	if s[a] > b {
		s[c] = 1
	} else {
		s[c] = 0
	}
	return s
}

func (s state) gtrr(a, b, c int) state {
	if s[a] > s[b] {
		s[c] = 1
	} else {
		s[c] = 0
	}
	return s
}

func (s state) eqir(a, b, c int) state {
	if a == s[b] {
		s[c] = 1
	} else {
		s[c] = 0
	}
	return s
}

func (s state) eqri(a, b, c int) state {
	if s[a] == b {
		s[c] = 1
	} else {
		s[c] = 0
	}
	return s
}

func (s state) eqrr(a, b, c int) state {
	if s[a] == s[b] {
		s[c] = 1
	} else {
		s[c] = 0
	}
	return s
}

type operation func(state, int, int, int) state

var ops = map[string]operation{
	"addr": state.addr, 
	"addi": state.addi, 
	"mulr": state.mulr, 
	"muli": state.muli,
	"banr": state.banr, 
	"bani": state.bani, 
	"borr": state.borr, 
	"bori": state.bori,
	"setr": state.setr, 
	"seti": state.seti,
	"gtir": state.gtir, 
	"gtri": state.gtri, 
	"gtrr": state.gtrr,
	"eqir": state.eqir, 
	"eqri": state.eqri, 
	"eqrr": state.eqrr,
}

type instr struct {
	op operation
	a, b, c int
}

type machine struct {
	reg state
	ip, ipreg int
	program []instr
}

func (m *machine) run() {
	for m.ip >= 0 && m.ip < len(m.program) {
		//fmt.Printf("ip=%d reg=%v\n", m.ip, m.reg)
		m.reg[m.ipreg] = m.ip
		i := m.program[m.ip]
		m.reg = i.op(m.reg, i.a, i.b, i.c)
		m.ip = m.reg[m.ipreg]
		m.ip++
	}
}

func main() {
	var m machine
	exp.MustForEachLineIn("inputs/19.txt", func(line string) {
		if strings.HasPrefix(line, "#ip") {
			if _, err := fmt.Sscanf(line, "#ip %d", &m.ipreg); err != nil {
				log.Fatalf("Couldn't parse directive: %v", err)
			}
			return
		}
		
		var oc string
		var i instr
		if _, err := fmt.Sscanf(line, "%s %d %d %d", &oc, &i.a, &i.b, &i.c); err != nil {
			log.Fatalf("Couldn't parse instruction: %v", err)
		}
		op, ok := ops[oc]
		if !ok {
			log.Fatalf("Invalid opcode %q", oc)
		}
		i.op = op
		m.program = append(m.program, i)
	})
	
	m.run()
	
	fmt.Println(m.reg[0])
}

