package main

import (
	"bufio"
	"fmt"
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

func mmmap(r rune) rune {
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

	mem := make(map[int]uint64)
	var mb, mm uint64 // mask bits, mask mask :P
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
			mms := strings.Map(mmmap, maskc[1])
			mmi, err := strconv.ParseUint(mms, 2, 64)
			if err != nil {
				die("Couldn't parse %q: %v", mms, err)
			}
			mb, mm = mbi, mmi
		case len(memc) == 3:
			addr, err := strconv.Atoi(memc[1])
			if err != nil {
				die("Couldn't parse %q: %v", memc[1], err)
			}
			val, err := strconv.ParseUint(memc[2], 10, 64)
			if err != nil {
				die("Couldn't parse %q: %v", memc[2], err)
			}
			mem[addr] = val&mm | mb
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
