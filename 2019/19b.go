package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
)

const generate = false
const find = true

func main() {
	ovm := intcode.ReadProgram("inputs/19.txt")

	test := func(x, y int) bool {
		vm := ovm.Copy()
		in, out := make(chan int, 2), make(chan int, 1)
		in <- x
		in <- y
		close(in)
		vm.Run(in, out)
		return <-out == 1
	}

	bounds := image.Rect(0, 0, 3000, 3000)

	if generate {
		bw := color.Palette{color.Black, color.White}
		img := image.NewPaletted(bounds, bw)
		var imgmu sync.Mutex
		work := make(chan image.Rectangle, runtime.NumCPU())
		var wg sync.WaitGroup
		for i := 0; i < runtime.NumCPU(); i++ {
			wg.Add(1)
			go func() {
				for r := range work {
					imgmu.Lock()
					subimg := img.SubImage(r).(*image.Paletted)
					imgmu.Unlock()
					for x := r.Min.X; x < r.Max.X; x++ {
						for y := r.Min.Y; y < r.Max.Y; y++ {
							if test(x, y) {
								subimg.Set(x, y, color.White)
							}
						}
					}
				}
				wg.Done()
			}()
		}

		for x := 0; x < 3000; x += 100 {
			for y := 0; y < 3000; y += 100 {
				work <- image.Rect(x, y, x+100, y+100)
			}
		}
		close(work)

		wg.Wait()
		f, err := os.Create("inputs/19b.png")
		if err != nil {
			log.Fatalf("Couldn't create output image: %v", err)
		}
		defer f.Close()
		if err := png.Encode(f, img); err != nil {
			log.Fatalf("Couldn't encode output image: %v", err)
		}
		if err := f.Close(); err != nil {
			log.Fatalf("Couldn't close output image: %v", err)
		}
	}

	if find {
		f, err := os.Open("inputs/19b.png")
		if err != nil {
			log.Fatalf("Couldn't open image: %v", err)
		}
		defer f.Close()

		img, err := png.Decode(f)
		if err != nil {
			log.Fatalf("Couldn't decode image: %v", err)
		}
		at := func(x, y int) int {
			return int(img.(*image.Paletted).ColorIndexAt(x, y))
		}

		grid := make([][]int, bounds.Max.Y)
		for y := 0; y < bounds.Max.Y; y++ {
			grid[y] = make([]int, bounds.Max.X)
			if y == 0 {
				for x := 0; x < bounds.Max.X; x++ {
					if x == 0 {
						grid[y][x] = at(x, y)
						continue
					}
					grid[y][x] = grid[y][x-1] + at(x, y)
				}
				continue
			}
			for x := 0; x < bounds.Max.X; x++ {
				if x == 0 {
					grid[y][x] = grid[y-1][x] + at(x, y)
					continue
				}
				grid[y][x] = grid[y-1][x] + grid[y][x-1] - grid[y-1][x-1] + at(x, y)
				if x > 100 && y > 100 {
					if grid[y][x]-grid[y-100][x]-grid[y][x-100]+grid[y-100][x-100] == 10000 {
						fmt.Println(x-99, y-99)
						return
					}
				}
			}
		}
	}
}
