package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"drjosh.dev/exp"
)

var errInvalidOp = errors.New("invalid operation")

func isreg(a int) bool {
	return a >= 0 && a < 4
}

func isreg2(a, b int) bool {
	return a >= 0 && a < 4 && b >= 0 && b < 4
}

func isreg3(a, b, c int) bool {
	return a >= 0 && a < 4 && b >= 0 && b < 4 && c >= 0 && c < 4
}

type state [4]int

func (s state) addr(a, b, c int) (state, error) {
	if !isreg3(a, b, c) {
		return state{}, errInvalidOp
	}
	s[c] = s[a] + s[b]
	return s, nil
}

func (s state) addi(a, b, c int) (state, error) {
	if !isreg2(a, c) {
		return state{}, errInvalidOp
	}
	s[c] = s[a] + b
	return s, nil
}

func (s state) mulr(a, b, c int) (state, error) {
	if !isreg3(a, b, c) {
		return state{}, errInvalidOp
	}
	s[c] = s[a] * s[b]
	return s, nil
}

func (s state) muli(a, b, c int) (state, error) {
	if !isreg2(a, c) {
		return state{}, errInvalidOp
	}
	s[c] = s[a] * b
	return s, nil
}

func (s state) banr(a, b, c int) (state, error) {
	if !isreg3(a, b, c) {
		return state{}, errInvalidOp
	}
	s[c] = s[a] & s[b]
	return s, nil
}

func (s state) bani(a, b, c int) (state, error) {
	if !isreg2(a, c) {
		return state{}, errInvalidOp
	}
	s[c] = s[a] & b
	return s, nil
}

func (s state) borr(a, b, c int) (state, error) {
	if !isreg3(a, b, c) {
		return state{}, errInvalidOp
	}
	s[c] = s[a] | s[b]
	return s, nil
}

func (s state) bori(a, b, c int) (state, error) {
	if !isreg2(a, c) {
		return state{}, errInvalidOp
	}
	s[c] = s[a] | b
	return s, nil
}

func (s state) setr(a, _, c int) (state, error) {
	if !isreg2(a, c) {
		return state{}, errInvalidOp
	}
	s[c] = s[a]
	return s, nil
}

func (s state) seti(a, _, c int) (state, error) {
	if !isreg(c) {
		return state{}, errInvalidOp
	}
	s[c] = a
	return s, nil
}

func (s state) gtir(a, b, c int) (state, error) {
	if !isreg2(b, c) {
		return state{}, errInvalidOp
	}
	if a > s[b] {
		s[c] = 1
	} else {
		s[c] = 0
	}
	return s, nil
}

func (s state) gtri(a, b, c int) (state, error) {
	if !isreg2(a, c) {
		return state{}, errInvalidOp
	}
	if s[a] > b {
		s[c] = 1
	} else {
		s[c] = 0
	}
	return s, nil
}

func (s state) gtrr(a, b, c int) (state, error) {
	if !isreg3(a, b, c) {
		return state{}, errInvalidOp
	}
	if s[a] > s[b] {
		s[c] = 1
	} else {
		s[c] = 0
	}
	return s, nil
}

func (s state) eqir(a, b, c int) (state, error) {
	if !isreg2(b, c) {
		return state{}, errInvalidOp
	}
	if a == s[b] {
		s[c] = 1
	} else {
		s[c] = 0
	}
	return s, nil
}

func (s state) eqri(a, b, c int) (state, error) {
	if !isreg2(a, c) {
		return state{}, errInvalidOp
	}
	if s[a] == b {
		s[c] = 1
	} else {
		s[c] = 0
	}
	return s, nil
}

func (s state) eqrr(a, b, c int) (state, error) {
	if !isreg3(a, b, c) {
		return state{}, errInvalidOp
	}
	if s[a] == s[b] {
		s[c] = 1
	} else {
		s[c] = 0
	}
	return s, nil
}

type operation func(state, int, int, int) (state, error)

var ops = []operation{
	state.addr, state.addi, state.mulr, state.muli,
	state.banr, state.bani, state.borr, state.bori,
	state.setr, state.seti,
	state.gtir, state.gtri, state.gtrr,
	state.eqir, state.eqri, state.eqrr,
}

func main() {
	/*
		possible := make([]uint16, 16)
		for i := range possible {
			possible[i] = ^uint16(0)
		}
		// possible[n] = bitset of ops that opcode n could still represent
	*/

	wantBefore := true
	var before state
	var opcode, a, b, c int
	examples := 0
	count := 0
	exp.MustForEachLineIn("inputs/16.txt", func(line string) {
		switch {
		case strings.TrimSpace(line) == "":
			break

		case strings.HasPrefix(line, "Before:"):
			if !wantBefore {
				log.Fatal("Unexpected Before line")
			}
			wantBefore = false
			if _, err := fmt.Sscanf(line, "Before: [%d, %d, %d, %d]", &before[0], &before[1], &before[2], &before[3]); err != nil {
				log.Fatalf("Couldn't scan line %q: %v", line, err)
			}

		case strings.HasPrefix(line, "After:"):
			if wantBefore {
				log.Fatal("Unexpected After line")
			}
			wantBefore = true
			var want state
			if _, err := fmt.Sscanf(line, "After: [%d, %d, %d, %d]", &want[0], &want[1], &want[2], &want[3]); err != nil {
				log.Fatalf("Couldn't scan line %q: %v", line, err)
			}

			examples++

			pass := 0
			for _, op := range ops {
				got, err := op(before, a, b, c)
				if err == nil && got == want {
					pass++
				}
				// else { possible[opcode] &^= 1 << i }
			}
			if pass >= 3 {
				count++
			}

		case !wantBefore:
			if _, err := fmt.Sscanf(line, "%d %d %d %d", &opcode, &a, &b, &c); err != nil {
				log.Fatalf("Couldn't scan line %q: %v", line, err)
			}

		default:
			// Reached the test program
			fmt.Println(count, "out of", examples)
			os.Exit(0)
		}
	})

}
