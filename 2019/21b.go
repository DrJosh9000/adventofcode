package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

var (
	opcodes = []string{"AND", "OR", "NOT"}
	rwregs  = []string{"T", "J"}
	rregs   = []string{"T", "J", "A", "B", "C", "D", "E", "F", "G", "H", "I"}

	ovm *intcode.VM
)

type inst struct {
	opcode, in, out int
}

func randInst() inst {
	// return an instruction that does something.
	for {
		i := inst{rand.Intn(3), rand.Intn(11), rand.Intn(2)}
		if i.opcode == 2 || i.in != i.out {
			return i
		}
	}
}

type prog struct {
	i [15]inst
	l int
}

func (p prog) String() string {
	var b strings.Builder
	for i := 0; i < p.l; i++ {
		in := p.i[i]
		fmt.Fprintf(&b, "%s %s %s\n", opcodes[in.opcode], rregs[in.in], rwregs[in.out])
	}
	fmt.Fprintf(&b, "RUN\n\n")
	return b.String()
}

func (p *prog) eval() int {
	input := p.String()
	inpos := 0
	vm := ovm.Copy()
	in, out := make(chan int), make(chan int)
	go vm.Run(in, out)
vmLoop:
	for {
		select {
		case x, ok := <-out:
			if !ok {
				break vmLoop
			}
			if x < 128 {
				break
			}
			fmt.Println("Winner!")
			fmt.Println("Program:")
			fmt.Print(p)
			fmt.Println("Score:", vm.CycleCount())
			fmt.Println("Damage:", x)
			os.Exit(0)
			break vmLoop
		case in <- int(input[inpos]):
			inpos++
		}
	}
	return vm.CycleCount()
}

func (p *prog) mutate() {
	switch rand.Intn(4) {
	case 0: // change one instruction
		i := rand.Intn(p.l)
		p.i[i] = randInst()
	case 1: // insert an instruction
		if p.l >= 15 {
			break
		}
		p.i[p.l] = randInst()
		i := rand.Intn(p.l)
		p.i[i], p.i[p.l] = p.i[p.l], p.i[i]
		p.l++
	case 2: // delete an instruction
		if p.l <= 1 {
			break
		}
		i := rand.Intn(p.l)
		for j := i; j < p.l-1; j++ {
			p.i[j] = p.i[j+1]
		}
		p.l--
		p.i[p.l] = inst{}
	case 3: // swap two instructions
		i, j := rand.Intn(p.l), rand.Intn(p.l)
		p.i[i], p.i[j] = p.i[j], p.i[i]
	}
}

func main() {
	ovm = intcode.ReadProgram("inputs/21.txt")

	const N = 500

	type progscore struct {
		p prog
		s int
	}
	var ranked []progscore
	pool := make(map[prog]int)
	for i := 0; len(pool) < N; i++ {
		var p prog
		p.l = (i % 14) + 1
		for j := 0; j < p.l; j++ {
			p.i[j] = randInst()
		}
		if _, exists := pool[p]; exists {
			continue
		}
		s := p.eval()
		pool[p] = s
		ranked = append(ranked, progscore{p, s})
	}

	itcount := 0
	for {
		sort.Slice(ranked, func(i, j int) bool {
			return ranked[i].s > ranked[j].s
		})
		log.Printf("%d iterations: best score %d", itcount, ranked[0].s)

		var mu sync.RWMutex
		var wg sync.WaitGroup
		wg.Add(runtime.NumCPU())
		work := make(chan prog)
		for i := 0; i < runtime.NumCPU(); i++ {
			go func() {
				for p := range work {
					known := true
					for known {
						p.mutate()
						mu.RLock()
						_, known = pool[p]
						mu.RUnlock()
					}
					s := p.eval()
					mu.Lock()
					pool[p] = s
					ranked = append(ranked, progscore{p, s})
					mu.Unlock()
				}
				wg.Done()
			}()
		}
		for _, ps := range ranked[:N] {
			work <- ps.p
		}
		close(work)
		wg.Wait()

		itcount++
	}
}
