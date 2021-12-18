package stuff

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTranspose(t *testing.T) {
	g := Matrix[int]{
		[]int{0, 1, 2, 3},
		[]int{3, 4, 5, 6},
		[]int{6, 7, 8, 9},
	}
	want := Matrix[int]{
		[]int{0, 3, 6},
		[]int{1, 4, 7},
		[]int{2, 5, 8},
		[]int{3, 6, 9},
	}
	got := g.Transpose()
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("g.Transpose diff:\n%s", diff)
	}
}

func TestFlipHorizontal(t *testing.T) {
	g := Matrix[int]{
		[]int{0, 1, 2, 3},
		[]int{3, 4, 5, 6},
		[]int{6, 7, 8, 9},
	}
	want := Matrix[int]{
		[]int{3, 2, 1, 0},
		[]int{6, 5, 4, 3},
		[]int{9, 8, 7, 6},
	}
	got := g.FlipHorizontal()
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("g.FlipHorizontal diff:\n%s", diff)
	}
}

func TestFlipVerticalal(t *testing.T) {
	g := Matrix[int]{
		[]int{0, 1, 2, 3},
		[]int{3, 4, 5, 6},
		[]int{6, 7, 8, 9},
	}
	want := Matrix[int]{
		[]int{6, 7, 8, 9},
		[]int{3, 4, 5, 6},
		[]int{0, 1, 2, 3},
	}
	got := g.FlipVertical()
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("g.FlipVertical diff:\n%s", diff)
	}
}

func TestRotate(t *testing.T) {
	g := Matrix[int]{
		[]int{0, 1, 2, 3},
		[]int{3, 4, 5, 6},
		[]int{6, 7, 8, 9},
	}
	want := Matrix[int]{
		[]int{6, 3, 0},
		[]int{7, 4, 1},
		[]int{8, 5, 2},
		[]int{9, 6, 3},
	}
	got := g.Rotate()
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("g.Rotate diff:\n%s", diff)
	}
}
