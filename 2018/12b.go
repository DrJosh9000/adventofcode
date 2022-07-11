package main

import (
	"fmt"
	"strings"

	"github.com/DrJosh9000/exp"
)

func main() {
	var state string
	rules := make(map[string]string)
	exp.MustForEachLineIn("inputs/12.txt", func(line string) {
		if strings.HasPrefix(line, "initial state: ") {
			state = strings.TrimPrefix(line, "initial state: ") 
			return
		}
		parts := strings.Split(line, " => ")
		if len(parts) != 2 {
			return
		}
		rules[parts[0]] = parts[1]
	})
	//fmt.Println(state)
	
	pad := strings.Repeat(".", 1000)
	state = pad + state + pad
	left := len(pad)
	
	sum := func() int {
		s := 0
		for i, c := range state {
			if c == '#' {
				s += i - left
			}
		}
		return s
	}
	
	last := sum()
	ldiff := 0
	for g := 1; ; g++ {
		var sb strings.Builder
		sb.WriteString("..")
		for i := 2; i < len(state)-2; i++ {
			o := rules[state[i-2:i+3]]
			if o == "" {
				o = state[i:i+1]
			}
			sb.WriteString(o)
		}
		sb.WriteString("..")
		state = sb.String()
		//fmt.Println(state)
		s := sum()
		d := s - last
		if d == ldiff {
			fmt.Println((50_000_000_000 - g) * d + s)
			return
		}
		last, ldiff = s, d
	}
}