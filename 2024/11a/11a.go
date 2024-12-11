package main

import (
	_ "embed"
	"fmt"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

//go:embed inputs/11.txt
var input string

func main() {
	line := algo.ListFromSlice(algo.Map(strings.Fields(input), exp.MustAtoi))
	start := line[0]
	for range 25 {
		p := start
		for {
			func() {
				if p.Value == 0 {
					p.Value = 1
					return
				}

				if l := ln10(p.Value); l%2 == 0 {
					left := p.Value
					rf := 1
					right := 0
					for range l / 2 {
						right += rf * (left % 10)
						rf *= 10
						left /= 10
					}
					n := &algo.ListNode[int]{Value: left}
					n.InsertBefore(p)
					if p == start {
						start = n
					}
					p.Value = right
					return
				}

				p.Value *= 2024
			}()

			p = p.Next
			if p == start {
				break
			}
		}

		// p = start
		// for {
		// 	fmt.Printf(" %d", p.Value)
		// 	p = p.Next
		// 	if p == start {
		// 		break
		// 	}
		// }
		// fmt.Println()
	}

	c := 1
	for p := start.Next; p != start; p = p.Next {
		c++
	}

	fmt.Println(c)
}

func ln10(n int) int {
	c := 0
	for n > 0 {
		n /= 10
		c++
	}
	return c
}
