package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"slices"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2023
// Day 24, part b

const inputPath = "2023/inputs/24.txt"

var big0 = big.NewInt(0)

func main() {
	lines := exp.MustReadLines(inputPath)
	type stone struct {
		p, v  algo.Vec3[int]
		value *big.Int
	}

	stones := algo.Map(lines, func(line string) (s stone) {
		exp.MustSscanf(line, "%d, %d, %d @ %d, %d, %d", &s.p[0], &s.p[1], &s.p[2], &s.v[0], &s.v[1], &s.v[2])
		return s
	})

	dist := func(a, b stone) *big.Int {
		//   ((apx+avx*t) - (bpx+bvx*t))^2 + ...
		// = ((apx-bpx) + (avx-bvx)*t)^2 + ...
		// = (dpx^2 + 2 dpx dvx t + dvx^2 t^2) +
		//   (dpy^2 + 2 dpy dvy t + dvy^2 t^2) +
		//   (dpz^2 + 2 dpz dvz t + dvz^2 t^2)
		// min at t = -(dp.dv) / (dv.dv)
		dp := bigVec3FromVec3Int(b.p.Sub(a.p))
		dv := bigVec3FromVec3Int(b.v.Sub(a.v))
		dvv := dv.dot(dv)
		if dvv.Cmp(big0) == 0 {
			return nil
		}
		dpv := dp.dot(dv)

		// So dist(t) = dp.dp + 2 dp.dv t + dv.dv t^2
		//            = dp.dp - 2(dp.dv)^2 / dv.dv + (dv.dv)(dp.dv)^2 / (dv.dv)^2
		//            = dp.dp - 2(dp.dv)^2 / dv.dv + (dp.dv)^2 / (dv.dv)
		//            = dp.dp - (dp.dv)^2 / dv.dv
		//t := -dpv / dvv
		// if t < 0 {
		// 	return math.Inf(1)
		// }
		y := dp.dot(dp)
		dpv.Mul(dpv, dpv)
		dpv.Div(dpv, dvv)
		y.Sub(y, dpv)
		return y
	}

	eval := func(a stone) *big.Int {
		sum := big.NewInt(0)
		for _, s := range stones {
			sum.Add(sum, dist(a, s))
		}
		return sum
	}

	const trl = 1_000_000_000_000
	ord := []int{1, 100, 10000, 1000000, 1_000_000_000, trl}

	pool := []stone{{
		p: algo.Vec3[int]{300 * trl, 250 * trl, 120 * trl},
		v: algo.Vec3[int]{-125, 25, 272},
	}}
	pool[0].value = eval(pool[0])
	best := pool[0].value
	fmt.Println("initial best:", best)
	for {
		lim := min(len(pool)/2+1, 100)
		pool = pool[:lim]
		for _, a := range pool {
			for i := range a.p {
				for _, m := range ord {
					b := a
					b.p[i] += m * rands()
					b.value = eval(b)
					if b.value != nil {
						pool = append(pool, b)
					}

					b = a
					b.p[i] += rand.Intn(2*m) - m
					// b.v[i] += rand.Float64()*100 - 50
					b.value = eval(b)
					if b.value != nil {
						pool = append(pool, b)
					}
				}
				// b := a
				// b.v[i] += s
				// b.value = eval(b)
				// pool = append(pool, b)
			}

		}

		slices.SortFunc(pool, func(a, b stone) int {
			return a.value.Cmp(b.value)
		})

		if pool[0].value.Cmp(best) < 0 {
			best = pool[0].value
			fmt.Println(pool[0].p, pool[0].v, pool[0].value.String())
		}

		if pool[0].value.Cmp(big0) == 0 {
			fmt.Println(pool[0].p)
			fmt.Println(pool[0].p.L1())
			return
		}
	}
}

func rands() int {
	r := rand.Intn(2)
	if r == 0 {
		return -1
	}
	return 1
}

type bigVec3 [3]*big.Int

// func newBigVec3() bigVec3 {
// 	return bigVec3{big.NewInt(0), big.NewInt(0), big.NewInt(0)}
// }

func bigVec3FromVec3Int(x algo.Vec3[int]) bigVec3 {
	return bigVec3{big.NewInt(int64(x[0])), big.NewInt(int64(x[1])), big.NewInt(int64(x[2]))}
}

// func (x bigVec3) add(y bigVec3) bigVec3 {
// 	z := newBigVec3()
// 	z[0].Add(x[0], y[0])
// 	z[1].Add(x[1], y[1])
// 	z[2].Add(x[2], y[2])
// 	return z
// }

// func (x bigVec3) sub(y bigVec3) bigVec3 {
// 	z := newBigVec3()
// 	z[0].Sub(x[0], y[0])
// 	z[1].Sub(x[1], y[1])
// 	z[2].Sub(x[2], y[2])
// 	return z
// }

func (x bigVec3) dot(y bigVec3) *big.Int {
	z := big.NewInt(0)
	t := big.NewInt(0)
	t.Mul(x[0], y[0])
	z.Add(z, t)
	t.Mul(x[1], y[1])
	z.Add(z, t)
	t.Mul(x[2], y[2])
	z.Add(z, t)
	return z
}
