package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
)

// Advent of Code 2016
// Day 14, part a

func main() {
	salt := append(make([]byte, 0, 64), []byte(os.Args[1])...)

	type hash struct {
		i  uint64
		t3 byte
		t5 uint
	}
	reps := func(h []byte) (t3 byte, t5 uint) {
		t3 = 0xff
		l, r := byte(0xff), 0
		for _, c := range h {
			for _, d := range []byte{c >> 4, c & 0xf} {
				if d != l {
					l, r = d, 0
				}
				r++
				// "Only consider the first such triplet in a hash"
				if r == 3 && t3 == 0xff {
					t3 = d
				}
				if r >= 5 {
					t5 |= 1 << d
				}
			}
		}
		return t3, t5
	}

	ch := make(chan hash, 64)

	go func() {
		for i := uint64(0); ; i++ {
			h := md5.Sum(strconv.AppendUint(salt, i, 10))
			t3, t5 := reps(h[:])
			if t3 != 0xff {
				ch <- hash{i, t3, t5}
			}
		}
	}()

	q := []hash{<-ch}
	keys := 0
	for {
		h := q[0]
		q = q[1:]
		for len(q) == 0 || q[len(q)-1].i <= h.i+1000 {
			q = append(q, <-ch)
		}
		for _, k := range q {
			if k.i > h.i+1000 {
				break
			}
			if (1<<h.t3)&k.t5 != 0 {
				keys++
				if keys >= 64 {
					fmt.Println(h.i)
					return
				}
				break
			}
		}
	}
}
