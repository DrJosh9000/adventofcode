package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	maskRE = regexp.MustCompile(`^mask = ([01X]+)$`)
	memRE  = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m+"\n", p...)
	os.Exit(1)
}

func mbmap(r rune) rune {
	if r == 'X' {
		return '0'
	}
	return r
}

func mfmap(r rune) rune {
	if r == 'X' {
		return '1'
	}
	return '0'
}

func main() {
	f, err := os.Open("input.14")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	// Seems there's no more than 9 X-bits in each mask, so do it the direct way.

	mem := make(map[uint64]uint64)
	var mb, mf uint64 // or-mask bits, float bits, and-mask bits
	var oc int
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := sc.Text()
		maskc := maskRE.FindStringSubmatch(m)
		memc := memRE.FindStringSubmatch(m)
		switch {
		case len(maskc) == 2:
			mbs := strings.Map(mbmap, maskc[1])
			mbi, err := strconv.ParseUint(mbs, 2, 64)
			if err != nil {
				die("Couldn't parse %q: %v", mbs, err)
			}
			mfs := strings.Map(mfmap, maskc[1])
			mfi, err := strconv.ParseUint(mfs, 2, 64)
			if err != nil {
				die("Couldn't parse %q: %v", mfs, err)
			}
			mb, mf, oc = mbi, mfi, bits.OnesCount64(mfi)
			//fmt.Printf("mask = %s\nmb   = %036b\nmf   = %036b\n", maskc[1], mb, mf)
		case len(memc) == 3:
			addr, err := strconv.ParseUint(memc[1], 10, 64)
			if err != nil {
				die("Couldn't parse %q: %v", memc[1], err)
			}
			val, err := strconv.ParseUint(memc[2], 10, 64)
			if err != nil {
				die("Couldn't parse %q: %v", memc[2], err)
			}
			//fmt.Printf("addr        = %036b\n", addr)
			addr &^= mf
			//fmt.Printf("addr&^mf    = %036b\n", addr)
			addr |= mb
			//fmt.Printf("addr&^mf|mb = %036b\n", addr)
			for n := 0; n < (1 << oc); n++ {
				fl := uint64(0)
				pos := 0
				for j := 0; j < oc; j++ {
					for mf&(1<<pos) == 0 {
						pos++
					}
					if n&(1<<j) == 0 {
						pos++
						continue
					}
					fl |= 1 << pos
					pos++
				}
				//fmt.Printf("n, fl, addr|fl = %09b, %036b, %036b\n", n, fl, addr|fl)
				mem[addr|fl] = val
			}
		default:
			die("Unmatched line %q", m)
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}

	var sum uint64
	for _, c := range mem {
		sum += c
	}
	fmt.Println(sum)
}
