package main

import (
	"container/heap"
	"fmt"
	"image"
	"log"
	"math"
	"os"
)

func main() {
	input, err := os.ReadFile("inputs/23.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}

	var initial state
	i, j := 0, 0
	for _, b := range input {
		if !(b == '.' || b == 'A' || b == 'B' || b == 'C' || b == 'D') {
			continue
		}
		i++
		if i%2 == 0 && i >= 2 && i <= 8 {
			continue
		}
		initial[j] = b
		j++
	}
	fmt.Println("initial state:", initial)

	energy := map[state]int{initial: 0}
	visited := make(map[state]struct{})
	prev := make(map[state]state)
	queue := &priqueue{item{initial, 0}}

	relax := func(s state, i, j int) {
		ns := s.swap(i, j)
		if _, vis := visited[ns]; vis {
			return
		}
		e, exists := energy[ns]
		if !exists {
			e = math.MaxInt
		}
		if ne := energy[s] + cost[s[i]]*dist[i][j]; ne < e {
			energy[ns] = ne
			heap.Push(queue, item{s: ns, v: ne})
			prev[ns] = s
		}
	}

	for queue.Len() > 0 {
		s := heap.Pop(queue).(item).s
		//fmt.Println("visiting", s, "with energy", min)
		if s == target {
			// found it
			break
		}
		visited[s] = struct{}{}
	stateLoop:
		for i, b := range s {
			// don't try to move an empty cell
			if b == '.' {
				continue
			}
			if i <= 6 {
				// is target room still occupied by mismatching pods?
				if r0, r1 := s[room[b][0]], s[room[b][1]]; r0 != b && r0 != '.' && r1 != b && r1 != '.' {
					continue
				}
				// make the target the deepest free cell in the room
				targ := room[b][1]
				if s[targ] != '.' {
					targ = room[b][0]
				}
				if s[targ] != '.' {
					continue
				}
				// we're in the hall, try to move into target via the nearest
				// preroom.
				if i < preroom[b][0] {
					for j := i + 1; j <= preroom[b][0]; j++ {
						if s[j] != '.' {
							// blocked
							continue stateLoop
						}
					}
				}
				if i > preroom[b][1] {
					for j := preroom[b][1]; j < i; j++ {
						if s[j] != '.' {
							// blocked
							continue stateLoop
						}
					}
				}
				relax(s, i, targ)

			} else {
				// we're in a room; try to move out into the hall
				h := i - 6
				if i >= 11 {
					// we're in the bottom of a room, first check the cell above
					if s[i-4] != '.' {
						continue
					}
					h -= 4
				}
				// compute a state for each unobstructed cell to the left of the room in the hall
				for j := h; j >= 0; j-- {
					if s[j] != '.' {
						break
					}
					relax(s, i, j)
				}
				// compute a state for each unobstructed cell to the right of the room in the hall
				for j := h + 1; j <= 6; j++ {
					if s[j] != '.' {
						break
					}
					relax(s, i, j)
				}
			}

		}
	}
	//fmt.Println(energy)
	fmt.Println("traceback:")
	for t := target; t != (state{}); t = prev[t] {
		fmt.Println(t, energy[t])
	}
	fmt.Println(energy[target])
}

type state [15]byte

func (s state) String() string { return string(s[:]) }

func (s state) swap(i, j int) state {
	s[i], s[j] = s[j], s[i]
	return s
}

var target = state{
	'.', '.', '.', '.', '.', '.', '.', // 0 .. 6
	'A', 'B', 'C', 'D', //  7 .. 10
	'A', 'B', 'C', 'D', // 11 .. 14
}

var cost = map[byte]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}

var room = map[byte][2]int{
	'A': {7, 11},
	'B': {8, 12},
	'C': {9, 13},
	'D': {10, 14},
}

var preroom = map[byte][2]int{
	'A': {1, 2},
	'B': {2, 3},
	'C': {3, 4},
	'D': {4, 5},
}

var dist = make([][]int, 15)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func init() {
	for i := range dist {
		dist[i] = make([]int, 15)
	}
	cells := []image.Point{
		{0, 0}, {1, 0}, {3, 0}, {5, 0}, {7, 0}, {9, 0}, {10, 0},
		{2, 1}, {4, 1}, {6, 1}, {8, 1},
		{2, 2}, {4, 2}, {6, 2}, {8, 2},
	}
	for i, c := range cells {
		for j, d := range cells {
			dist[i][j] = abs(c.X-d.X) + abs(c.Y-d.Y)
		}
	}
}

type item struct {
	s state
	v int
}

type priqueue []item

func (pq priqueue) Len() int            { return len(pq) }
func (pq priqueue) Less(i, j int) bool  { return pq[i].v < pq[j].v }
func (pq priqueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priqueue) Push(x interface{}) { *pq = append(*pq, x.(item)) }
func (pq *priqueue) Pop() interface{} {
	n1 := len(*pq) - 1
	i := (*pq)[n1]
	*pq = (*pq)[0:n1]
	return i
}
