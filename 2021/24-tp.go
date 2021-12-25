package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var evals uint64

func main() {
	fmt.Println("starting search")
	start := time.Now()
	go func() {
		for range time.Tick(time.Minute) {
			e := atomic.LoadUint64(&evals)
			fmt.Println("evals/sec:", float64(e)/time.Since(start).Seconds())
		}
	}()
	var wg sync.WaitGroup
	for i := 9; i >= 1; i-- {
		i := i
		for j := 9; j >= 1; j-- {
			j := j
			wg.Add(1)
			go func() {
				search(append(make([]int, 0, 14), i, j))
				wg.Done()
			}()
		}
	}
	wg.Wait()
}

func search(in []int) {
	if len(in) == 14 {
		if eval(in) {
			fmt.Println(in)
		}
		return
	}
	for i := 9; i >= 1; i-- {
		search(append(in, i))
	}
}

func eval(in []int) bool {
	x, z := 0, 0
	z += in[0] + 14 // z = in[0]+14
	z *= 26         // z = (in[0]+14)*26
	z += in[1] + 8  // z = (in[0]+14)*26 + (in[1]+8)
	z *= 26         // z = (in[0]+14)*26^2 + (in[1]+8)*26
	z += in[2] + 4  // z = (in[0]+14)*26^2 + (in[1]+8)*26 + (in[2]+4)
	z *= 26         // z = (in[0]+14)*26^3 + (in[1]+8)*26^2 + (in[2]+4)*26
	z += in[3] + 10 // z = (in[0]+14)*26^3 + (in[1]+8)*26^2 + (in[2]+4)*26 + (in[3]+10)
	x = z%26 - 3    // x = (in[3]+10)-3 = in[3]+7
	z /= 26         // z = (in[0]+14)*26^2 + (in[1]+8)*26 + (in[2]+4)
	if x == in[4] { // in[3]+7 == in[4]
		x = 0
	} else {
		x = 1
	}
	z *= 25*x + 1         // x==0: nop
	z += (in[4] + 14) * x // x==0: nop
	x = z%26 - 4          // x = (in[2]+4)-4 = in[2]
	z /= 26               // z = (in[0]+14)*26 + (in[1]+8)
	if x == in[5] {       // in[2] == in[5]
		x = 0
	} else {
		x = 1
	}
	z *= 25*x + 1         // x==0: nop
	z += (in[5] + 10) * x // x==0: nop
	z *= 26               // z = (in[0]+14)*26^2 + (in[1]+8)*26
	z += in[6] + 4        // z = (in[0]+14)*26^2 + (in[1]+8)*26 + (in[6]+4)
	x = z%26 - 8          // x = (in[6]+4)-8 = in[6]-4
	z /= 26               // z = (in[0]+14)*26 + (in[1]+8)
	if x == in[7] {       // in[6]-4 == in[7]
		x = 0
	} else {
		x = 1
	}
	z *= 25*x + 1         // x==0: nop
	z += (in[7] + 14) * x // x==0: nop
	x = z%26 - 3          // x = (in[1]+8)-3 = in[1]+5
	z /= 26               // z = in[0]+14
	if x == in[8] {       // in[1]+5 == in[8]
		x = 0
	} else {
		x = 1
	}
	z *= 25*x + 1        // x==0: nop
	z += (in[8] + 1) * x // x==0: nop
	x = z%26 - 12        // x = (in[0]+14)-12 = in[0]+2
	z /= 26              // z = 0
	if x == in[9] {      // in[0]+2 == in[9]
		x = 0
	} else {
		x = 1
	}
	z *= 25*x + 1        // x==0: nop
	z += (in[9] + 6) * x // x==0: nop
	z *= 26              // z = 0
	z += in[10]          // z = in[10]
	x = z%26 - 6         // x = in[10]-6
	z /= 26              // z = 0
	if x == in[11] {     // in[10]-6 == in[11]
		x = 0
	} else {
		x = 1
	}
	z *= 25*x + 1         // x==0: nop
	z += (in[11] + 9) * x // x==0: nop
	z *= 26               // z = 0
	z += in[12] + 13      // z = in[12]+13
	x = z%26 - 12         // x = (in[12]+13)-12 = in[12]+1
	z /= 26               // z = 0
	if x == in[13] {      // in[12]+1 == in[13]
		x = 0
	} else {
		x = 1
	}
	z *= 25*x + 1          // x==0: nop
	z += (in[13] + 12) * x // x==0: nop
	atomic.AddUint64(&evals, 1)
	return z == 0
}

// in[3]+7 == in[4]
// in[2] == in[5]
// in[6]-4 == in[7]
// in[1]+5 == in[8]
// in[0]+2 == in[9]
// in[10]-6 == in[11]
// in[12]+1 == in[13]

var big = []int{
	0:  7,
	1:  4,
	2:  9,
	3:  2,
	4:  9,
	5:  9,
	6:  9,
	7:  5,
	8:  9,
	9:  9,
	10: 9,
	11: 3,
	12: 8,
	13: 9,
}

// 74929995999389

var small = []int{
	0:  1,
	1:  1,
	2:  1,
	3:  1,
	4:  8,
	5:  1,
	6:  5,
	7:  1,
	8:  6,
	9:  3,
	10: 7,
	11: 1,
	12: 1,
	13: 2,
}

// 11118151637112
