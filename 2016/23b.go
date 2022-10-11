package main

import (
	"fmt"
)

// Advent of Code 2016
// Day 23, part b

func main() {
	//      a, b, c, d := 12, 0, 0, 0
	//      b = a
	//      b--     // b = a-1 = 11
	// l2:  d = a   // d = 12, 132, 1320, ...
	//      a = 0
	// l4:  c = b   // c = 11, 10
	// l5:  a++
	//      c--
	//      if c != 0 goto l5  // a += c; c = 0
	//      d--
	//      if d != 0 goto l4      // a = a*b (132, 120, 108, 96, 84, 72, 60, 48, 36, 24, 12)
	//      b--                    // b = 10...1
	//      c = b
	//      d = c
	// l13: d--
	//      c++
	//      if d != 0 goto l13     // c, d = 2*b, 0 (20...2)
	// T0:  tgl c | c++
	//      c = -16
	// T2:  ~goto l2~ | c = 1        // toggled when b = 1
	//      c = 78                 // only get here after b = 1
	// T4:  ~goto d~ | d = 70        // toggled when b = 2
	// l21: a++
	// T6:  ~d++~ | d--              // toggled when b = 3
	//      if d != 0 goto l21
	// T8:  ~c++~ | c--              // toggled when b = 4
	//      if c != 0 goto T4

	a := 12
	b := a - 1
	for b > 0 {
		a *= b
		b--
	}
	a += 78 * 70

	fmt.Println(a)
}
