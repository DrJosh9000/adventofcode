package main

import (
	"fmt"
	"math/bits"
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
	maxl, maxw := 0, 0
	weight := map[state]int{{}: 0}
	q := []state{{}}
	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		w := weight[s]
		l := bits.OnesCount64(s.used)
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
			lt := l + 1
			wt := w + d[0] + d[1]
			if lt > maxl || (lt == maxl && wt > maxw) {
				maxl = lt
				maxw = wt
			}
			weight[st] = wt
		}
	}
	fmt.Println(maxw)
}
