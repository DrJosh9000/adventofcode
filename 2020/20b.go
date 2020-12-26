package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

// Directions
const (
	Up = iota
	Down
	Left
	Right
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m+"\n", p...)
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

func bit2int(bmp []uint, n int) uint {
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
	// up, down, left, right
	f, r [4]uint

	// neighbors (in this tile's orientation)
	// up, down, left, right
	n [4]*tile

	bmp []uint
}

func (t *tile) edgenums() {
	t.f[Up] = t.bmp[0]
	t.f[Down] = t.bmp[9]
	t.f[Left] = bit2int(t.bmp, 9)
	t.f[Right] = bit2int(t.bmp, 0)
	for i, h := range t.f {
		t.r[i] = reverse(h)
	}
}

func (t *tile) rotate() {
	// left -> up -> right -> down -> left
	nbmp := make([]uint, len(t.bmp))
	for i := range t.bmp {
		nbmp[i] = reverse(bit2int(t.bmp, len(t.bmp)-i-1))
	}
	t.bmp = nbmp
	t.edgenums()
	// rotate neighbors too
	t.n[Up], t.n[Right], t.n[Down], t.n[Left] = t.n[Left], t.n[Up], t.n[Right], t.n[Down]
}

func (t *tile) flipH() {
	// left <-> right
	for i, row := range t.bmp {
		t.bmp[i] = reverse(row)
	}
	t.edgenums()
	t.n[Left], t.n[Right] = t.n[Right], t.n[Left]
}

func (t *tile) flipV() {
	// up <-> down
	nbmp := make([]uint, len(t.bmp))
	for i, row := range t.bmp {
		nbmp[len(t.bmp)-i-1] = row
	}
	t.bmp = nbmp
	t.edgenums()
	t.n[Up], t.n[Down] = t.n[Down], t.n[Up]
}

func (t *tile) String() string {
	sb := new(strings.Builder)
	fmt.Fprintf(sb, "Tile %d:\n", t.id)
	for _, row := range t.bmp {
		fmt.Fprintf(sb, "%010b\n", row)
	}
	fmt.Fprintf(sb, "fu, fd, fl, fr: %010b, %010b, %010b, %010b\n", t.f[0], t.f[1], t.f[2], t.f[3])
	fmt.Fprintf(sb, "ru, rd, rl, rr: %010b, %010b, %010b, %010b\n", t.r[0], t.r[1], t.r[2], t.r[3])
	return sb.String()
}

func search(picture []string, monster []string) int {
	count := 0
	for j, row := range picture[:len(picture)-len(monster)] {
		for i := range row[:len(row)-len(monster[0])] {
			nessie := true
		monsterCheck:
			for y, mrow := range monster {
				for x, r := range mrow {
					if r == ' ' {
						continue
					}
					if picture[j+y][i+x] == '0' {
						nessie = false
						break monsterCheck
					}
				}
			}
			if nessie {
				count++
			}
		}
	}
	return count
}

func rotate(m []string) []string {
	bs := make([]strings.Builder, len(m[0]))
	for i := range bs {
		for j := range m {
			bs[i].WriteByte(m[len(m)-j-1][i])
		}
	}
	r := make([]string, 0, len(bs))
	for _, b := range bs {
		r = append(r, b.String())
	}
	return r
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
			t = new(tile)
			t.id = atoui(m[5:9])
		case m == "":
			t.edgenums()
			tiles = append(tiles, t)
			//fmt.Println(ct)

			for _, n := range t.f {
				et[n] = append(et[n], t)
			}
			for _, n := range t.r {
				et[n] = append(et[n], t)
			}
		default:
			t.bmp = append(t.bmp, dothash2int(m))
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}

	// match tiles together
	for g, ts := range et {
		if len(ts) != 2 {
			continue
		}

		for i, h := range ts[0].f {
			if g == h {
				ts[0].n[i] = ts[1]
			}
		}
		for i, h := range ts[0].r {
			if g == h {
				ts[0].n[i] = ts[1]
			}
		}
		for i, h := range ts[1].f {
			if g == h {
				ts[1].n[i] = ts[0]
			}
		}
		for i, h := range ts[1].r {
			if g == h {
				ts[1].n[i] = ts[0]
			}
		}
	}

	// find a corner...
	var corner *tile
	for _, t := range tiles {
		c := 0
		for _, n := range t.n {
			if n != nil {
				c++
			}
		}
		if c == 2 {
			corner = t
			break
		}
	}

	// rotate corner until the neighbors are right and down
	for corner.n[Right] == nil || corner.n[Down] == nil {
		corner.rotate()
	}

	// set "corner" to be the tile at 0, 0.
	grid := make([][]*tile, 12)
	for j := range grid {
		grid[j] = make([]*tile, 12)
	}
	grid[0][0] = corner

	// Now arrange all the tiles in the grid.
	for j, row := range grid {
		for i := range row {
			if row[i] != nil {
				continue
			}
			if i == 0 {
				// This tile is the down neighbor of the same tile of the
				// previous row.
				row[i] = grid[j-1][i].n[Down]
				// Rotate this tile until n[Up] == tile above.
				for row[i].n[Up] != grid[j-1][i] {
					row[i].rotate()
				}
				// If the down edge id of the tile above doesn't match the up
				// edge id of this tile, flip this tile horizontally.
				if grid[j-1][i].f[Down] != row[i].f[Up] {
					row[i].flipH()
				}
				// Check: they match, right?
				if grid[j-1][i].f[Down] != row[i].f[Up] {
					die("Something went wrong rotating/flipping tile at %d, %d: %010b != %010b", j, i, grid[j-1][i].f[Down], row[i].f[Up])
				}
				continue
			}
			// This tile is the right neighbor of the previous tile
			row[i] = row[i-1].n[Right]
			// Rotate this tile until n[Left] == tile to the left.
			for row[i].n[Left] != row[i-1] {
				row[i].rotate()
			}
			// If the right edge id of the previous tile doesn't match the left
			// edge id of this tile, flip this tile vertically.
			if row[i-1].f[Right] != row[i].f[Left] {
				row[i].flipV()
			}
			// Check: they match, right?
			if row[i-1].f[Right] != row[i].f[Left] {
				die("Something went wrong rotating/flipping tile at %d, %d: %010b != %010b", j, i, row[i-1].f[Right], row[i].f[Left])
			}
			if j > 0 {
				// Check it matches the tile above too.
				if grid[j-1][i].f[Down] != row[i].f[Up] {
					die("Something went wrong rotating/flipping tile at %d, %d: %010b != %010b", j, i, grid[j-1][i].f[Down], row[i].f[Up])
				}
			}
		}
	}

	// The grid is complete! stitch the picture together. Also count "hashes"qq
	hashes := 0
	picture := make([]string, 0, 12*8)
	for _, row := range grid {
		bs := make([]strings.Builder, 8)
		for _, t := range row {
			for k, v := range t.bmp[1:9] {
				p := (v >> 1) & 0xff
				hashes += bits.OnesCount(p)
				fmt.Fprintf(&bs[k], "%08b", p)
			}
		}
		for _, b := range bs {
			picture = append(picture, b.String())
		}
	}

	seamonster := []string{
		`                  # `,
		`#    ##    ##    ###`,
		` #  #  #  #  #  #   `,
	}

	monsters := [][]string{seamonster}
	for i := 0; i < 3; i++ {
		monsters = append(monsters, rotate(monsters[i]))
	}
	// turns out flipping wasn't needed...
	for _, monster := range monsters {
		count := search(picture, monster)
		if count != 0 {
			fmt.Println(hashes - 15*count)
			return
		}
	}
}
