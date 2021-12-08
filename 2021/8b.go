package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("inputs/8.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	var lines [][]uint
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var line []uint
		for _, token := range strings.Fields(sc.Text()) {
			var n uint
			for _, r := range token {
				n += 1 << (r - 'a')
			}
			line = append(line, n)
		}
		lines = append(lines, line)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't sc.scan: %v", err)
	}

	sum := 0
	for _, line := range lines {
		// deduce permutation
		fwd := make(map[uint]int)
		rev := make([]uint, 10)
		for _, n := range line[:10] {
			switch bits.OnesCount(n) {
			case 2:
				fwd[n] = 1
				rev[1] = n
			case 3:
				fwd[n] = 7
				rev[7] = n
			case 4:
				fwd[n] = 4
				rev[4] = n
			case 7:
				fwd[n] = 8
				rev[8] = n
			}
		}
		// subtract 1 from 4; that's the top-left and middle segments
		tlmid := rev[4] - rev[1]
		// find 5, 3, and 2
		for _, n := range line[:10] {
			if bits.OnesCount(n) == 5 {
				switch {
				case n&tlmid == tlmid:
					fwd[n] = 5
					rev[5] = n
				case n&rev[1] == rev[1]:
					fwd[n] = 3
					rev[3] = n
				default:
					fwd[n] = 2
					rev[2] = n
				}
			}
		}

		// now just 6, 9, 0
		// 6 and 9 overlap tlmid, 0 does not
		// 9 overlaps 7, 6 does not
		for _, n := range line[:10] {
			if bits.OnesCount(n) == 6 {
				switch {
				case n&tlmid != tlmid:
					fwd[n] = 0
					rev[0] = n
				case n&rev[7] == rev[7]:
					fwd[n] = 9
					rev[9] = n
				default:
					fwd[n] = 6
					rev[6] = n
				}
			}
		}

		out := 0
		for _, n := range line[11:] {
			out *= 10
			out += fwd[n]
		}
		sum += out
	}

	fmt.Println(sum)
}
