package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"
)

func main() {
	var dominos [][2]int
	exp.MustForEachLineIn("inputs/24.txt", func(line string) {
		sp := strings.Split(line, "/")
		dominos = append(dominos, [2]int{
			exp.Must(strconv.Atoi(sp[0])),
			exp.Must(strconv.Atoi(sp[1])),
		})
	})

	type state struct {
		used uint64
		last int
	}
	maxw := 0
	weight := map[state]int{{}: 0}
	q := []state{{}}
	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		w := weight[s]
		for i, d := range dominos {
			if s.used&(1<<i) != 0 {
				continue
			}
			var p int
			switch s.last {
			case d[0]:
				p = d[1]
			case d[1]:
				p = d[0]
			default:
				continue
			}
			st := state{
				used: s.used | (1 << i),
				last: p,
			}
			if _, seen := weight[st]; seen {
				continue
			}
			q = append(q, st)
			wt := w + d[0] + d[1]
			weight[st] = wt
			if wt > maxw {
				maxw = wt
			}
		}
	}
	fmt.Println(maxw)
}
