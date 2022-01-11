package main

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

func main() {
	// Get a picture of the maze
	ovm := intcode.ReadProgram("inputs/17.txt")
	vm := ovm.Copy()
	out := make(chan int)
	go vm.Run(nil, out)

	orient := map[int]image.Point{
		'^': {0, -1},
		'v': {0, 1},
		'<': {-1, 0},
		'>': {1, 0},
	}
	var (
		maze [][]byte
		line []byte
		p, o image.Point
	)
	for x := range out {
		if x == '\n' {
			if len(line) > 0 {
				maze = append(maze, line)
			}
			line = nil
			continue
		}
		if x != '.' && x != '#' {
			// must be the robot
			p = image.Pt(len(line), len(maze))
			o = orient[x]
		}
		line = append(line, byte(x))

	}
	if len(line) > 0 {
		maze = append(maze, line)
	}

	fmt.Println("Robot at", p, "orientation", o)

	// Find a path through the maze
	left := func(q image.Point) image.Point { return image.Pt(q.Y, -q.X) }
	right := func(q image.Point) image.Point { return image.Pt(-q.Y, q.X) }
	bounds := image.Rect(0, 0, len(maze[0]), len(maze))
	onpath := func(q image.Point) bool {
		return q.In(bounds) && maze[q.Y][q.X] == '#'
	}

	var (
		path []string
		fwd  int
	)
pathfinder:
	for {
		switch {
		case onpath(p.Add(o)):
			fwd++
			p = p.Add(o)
		case onpath(p.Add(left(o))):
			if fwd > 0 {
				path = append(path, strconv.Itoa(fwd))
			}
			path = append(path, "L")
			o = left(o)
			fwd = 0
		case onpath(p.Add(right(o))):
			if fwd > 0 {
				path = append(path, strconv.Itoa(fwd))
			}
			path = append(path, "R")
			o = right(o)
			fwd = 0
		default:
			break pathfinder
		}
	}
	if fwd > 0 {
		path = append(path, strconv.Itoa(fwd))
	}
	fullpath := strings.Join(path, ",")
	//fmt.Println("Full path:", fullpath)

	// Compress the path by finding repeated subpaths.
	mp, sub := compress(fullpath, nil)
	fmt.Println("Main program:", mp)
	fmt.Println("Function A:", sub[0])
	fmt.Println("Function B:", sub[1])
	fmt.Println("Function C:", sub[2])
	input := strings.Join(append([]string{mp}, sub...), "\n") + "\nn\n"
	inpos := 0
	innum := func() int {
		if inpos >= len(input) {
			return 10
		}
		return int(input[inpos])
	}
	fmt.Printf("Full input: %q\n", input)

	// Run the program again
	vm = ovm.Copy()
	vm.Poke(0, 2)
	in, out := make(chan int), make(chan int)
	go vm.Run(in, out)
commLoop:
	for {
		select {
		case x, ok := <-out:
			if !ok {
				break commLoop
			}
			if x < 127 {
				fmt.Printf("%c", x)
			} else {
				fmt.Println(x)
			}
		case in <- innum():
			inpos++
		}
	}
}

var subpsym = []string{"A", "B", "C"}

func compress(in string, sub []string) (string, []string) {
	const fsym = "LR0123456789"
	if len(sub) == 3 {
		if len(in) > 20 {
			return "", nil
		}
		for _, s := range sub {
			if len(s) > 20 {
				return "", nil
			}
		}
		if strings.ContainsAny(in, fsym) {
			return "", nil
		}
		return in, sub
	}
	start := strings.IndexAny(in, fsym)
	if start < 0 {
		return "", nil
	}
	for end := start + 1; end < len(in); end++ {
		if in[end] == ',' {
			continue
		}
		nmp := strings.Replace(in, in[start:end+1], subpsym[len(sub)], -1)
		mp, sub := compress(nmp, append(sub, in[start:end+1]))
		if mp != "" && len(sub) > 0 {
			return mp, sub
		}
	}
	return "", nil
}
