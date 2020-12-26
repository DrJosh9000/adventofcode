package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m, p...)
	os.Exit(1)
}

func atoui(x string) uint {
	n, err := strconv.ParseUint(x, 10, 64)
	if err != nil {
		die("Couldn't parse %q: %v", x, err)
	}
	return uint(n)
}

func dothash2int(m string) uint {
	m = strings.Map(func(r rune) rune {
		switch r {
		case '.':
			return '0'
		case '#':
			return '1'
		}
		return r
	}, m)
	n, err := strconv.ParseUint(m, 2, 64)
	if err != nil {
		die("Couldn't parse %q: %v", m, err)
	}
	return uint(n)
}

func bit2int(bmp []uint, n uint) uint {
	r := uint(0)
	for _, d := range bmp {
		r <<= 1
		r += (d & (1 << n)) >> n
	}
	return r
}

func reverse(n uint) uint {
	return bits.Reverse(n) >> (bits.UintSize - 10)
}

type tile struct {
	id uint

	// cache of edge numbers
	fu, fd, fl, fr uint
	ru, rd, rl, rr uint

	neighs map[*tile]struct{}

	bmp []uint
}

func (t *tile) String() string {
	sb := new(strings.Builder)
	fmt.Fprintf(sb, "Tile %d:\n", t.id)
	for _, row := range t.bmp {
		fmt.Fprintf(sb, "%010b\n", row)
	}
	fmt.Fprintf(sb, "fu, fd, fl, fr: %010b, %010b, %010b, %010b\n", t.fu, t.fd, t.fl, t.fr)
	fmt.Fprintf(sb, "ru, rd, rl, rr: %010b, %010b, %010b, %010b\n", t.ru, t.rd, t.rl, t.rr)
	for neigh := range t.neighs {
		fmt.Fprintf(sb, "neigh id: %d\n", neigh.id)
	}
	return sb.String()
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	et := make(map[uint][]*tile)
	var tiles []*tile
	var t *tile
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := sc.Text()
		switch {
		case strings.HasPrefix(m, "Tile "):
			t = &tile{neighs: make(map[*tile]struct{})}
			t.id = atoui(m[5:9])
		case m == "":
			t.fu = t.bmp[0]
			t.fd = t.bmp[9]
			t.fl = bit2int(t.bmp, 9)
			t.fr = bit2int(t.bmp, 0)
			t.ru = reverse(t.fu)
			t.rd = reverse(t.fd)
			t.rl = reverse(t.fl)
			t.rr = reverse(t.fr)
			tiles = append(tiles, t)
			//fmt.Println(ct)

			et[t.fu] = append(et[t.fu], t)
			et[t.fd] = append(et[t.fd], t)
			et[t.fl] = append(et[t.fl], t)
			et[t.fr] = append(et[t.fr], t)
			et[t.ru] = append(et[t.ru], t)
			et[t.rd] = append(et[t.rd], t)
			et[t.rl] = append(et[t.rl], t)
			et[t.rr] = append(et[t.rr], t)
		default:
			t.bmp = append(t.bmp, dothash2int(m))
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}

	// Part A: find the corner pieces.
	// Solve the whole puzzle and find out?
	// hist of lengths of values in et: map[1:96 2:528]
	// So we're blessed with a bunch of unique edges.
	for _, ts := range et {
		if len(ts) != 2 {
			continue
		}
		ts[0].neighs[ts[1]] = struct{}{}
		ts[1].neighs[ts[0]] = struct{}{}
	}
	// Any tile with only two neighbors is a corner.
	prod := uint(1)
	for _, t := range tiles {
		if len(t.neighs) == 2 {
			prod *= t.id
		}
	}
	fmt.Println(prod)
}
