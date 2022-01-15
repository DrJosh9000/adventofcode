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
	rregs   = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "T", "J"}

	ovm *intcode.VM
)

type inst struct {
	opcode, in, out int
}

type prog struct {
	i []inst
	s int
}

func (p prog) String() string {
	var b strings.Builder
	for _, i := range p.i {
		fmt.Fprintf(&b, "%s %s %s\n", opcodes[i.opcode], rregs[i.in], rwregs[i.out])
	}
	fmt.Fprintf(&b, "RUN\n\n")
	return b.String()
}

func (p *prog) eval() {
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
			fmt.Println("Cycles:", vm.CycleCount())
			fmt.Println("Damage:", x)
			os.Exit(0)
			return
		case in <- int(input[inpos]):
			inpos++
		}
	}
	p.s = vm.CycleCount()
}

func (p *prog) mutate() {
	ni := append([]inst(nil), p.i...)
	n := len(ni)
	switch rand.Intn(3) {
	case 0: // change one instruction
		i := rand.Intn(n)
		ni[i] = inst{rand.Intn(3), rand.Intn(11), rand.Intn(2)}
	case 1: // insert an instruction
		if n >= 15 {
			break
		}
		ni = append(ni, inst{rand.Intn(3), rand.Intn(11), rand.Intn(2)})
		i := rand.Intn(n)
		ni[i], ni[n] = ni[n], ni[i]
	case 2: // delete an instruction
		if n <= 1 {
			break
		}
		i := rand.Intn(n)
		ni = append(ni[:i], ni[i+1:]...)
	case 3: // swap two instructions
		i, j := rand.Intn(n), rand.Intn(n)
		ni[i], ni[j] = ni[j], ni[i]
	}
	p.i = ni
}

func main() {
	ovm = intcode.ReadProgram("inputs/21.txt")

	const N = 1000

	pool := make([]prog, 0, N)
	for i := 0; i < N; i++ {
		var p prog
		for j := 0; j < (i%14)+1; j++ {
			p.i = append(p.i, inst{rand.Intn(3), rand.Intn(11), rand.Intn(2)})
		}
		p.eval()
		pool = append(pool, p)
	}

	itcount := 0
	for {
		sort.Slice(pool, func(i, j int) bool {
			return pool[i].s > pool[j].s
		})
		if itcount%10 == 0 {
			log.Printf("%d iterations: best score %d", itcount, pool[0].s)
		}
		pool = pool[:N/2]

		var mu sync.Mutex
		var wg sync.WaitGroup
		wg.Add(runtime.NumCPU())
		work := make(chan int)
		for i := 0; i < runtime.NumCPU(); i++ {
			go func() {
				for j := range work {
					p := pool[j]
					p.mutate()
					p.eval()
					mu.Lock()
					pool = append(pool, p)
					mu.Unlock()
				}
				wg.Done()
			}()
		}
		for i := 0; i < N/2; i++ {
			work <- i
		}
		close(work)
		wg.Wait()

		itcount++
	}
}
