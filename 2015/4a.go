package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
)

// Advent of Code 2015
// Day 4, part a

// Fuck, Santa...it's really gross of you to make me mine crapto-coins.
// Especially with a proof-of-work algorithm!

func main() {
	prefix := os.Args[1]
	N := runtime.NumCPU()
	best := uint64(math.MaxUint64)
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		i, N := uint64(i), uint64(N)
		go func() {
			defer wg.Done()
			buf := append(make([]byte, 0, 64), prefix...)
			for j := i; j < atomic.LoadUint64(&best); j += N {
				h := md5.Sum(strconv.AppendUint(buf, j, 10))
				if !(h[0] == 0 && h[1] == 0 && h[2] < 8) {
					continue
				}
				for {
					old := atomic.LoadUint64(&best)
					if j > old {
						return
					}
					if atomic.CompareAndSwapUint64(&best, old, j) {
						return
					}
				}
			}
		}()
	}
	wg.Wait()

	fmt.Println(best)
}
