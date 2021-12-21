// Package stuff is where I'm putting some useful tidbits from my AoC solutions.
package stuff

// Matrix is a "2D array" generic type.
type Matrix[T any] [][]T

// MakeMatrix makes a matrix of width w and height h.
func MakeMatrix[T any](w, h int) Matrix[T] {
	g := make(Matrix[T], h)
	for j := range g {
		g[j] = make([]T, w)
	}
	return g
}

// Size returns the width and height of the matrix. If the height is zero the
// width will also be zero.
func (g Matrix[T]) Size() (w, h int) {
	h = len(g)
	if h == 0 {
		return 0, 0
	} 
	return len(g[0]), h
}

// Fill fills the matrix with the value v.
func (g Matrix[T]) Fill(v T) {
	for _, row := range g {
		for i := range row {
			row[i] = v
		}
	}
}

// Transpose returns a new matrix reflected about the diagonal.
func (g Matrix[T]) Transpose() Matrix[T] {
	h, w := g.Size()
	ng := MakeMatrix[T](w, h)
	for j, row := range g {
		for i := range row {
			ng[i][j] = row[i]
		}
	}
	return ng
}


// FlipHorizontal returns a new matrix flipped horizontally (left becomes right).
func (g Matrix[T]) FlipHorizontal() Matrix[T] {
	w, h := g.Size()
	ng := MakeMatrix[T](w, h)
	for j, row := range g {
		for i := range row {
			ng[j][i] = row[w-i-1]
		}
	}
	return ng
}

// FlipVertical returns a new matrix flipped vertically (top becomes bottom).
func (g Matrix[T]) FlipVertical() Matrix[T] {
	w, h := g.Size()
	ng := MakeMatrix[T](w, h)
	for j, row := range g {
		for i := range row {
			ng[h-j-1][i] = row[i]
		}
	}
	return ng
}

// Rotate returns a new matrix rotated clockwise by 90 degrees.
func (g Matrix[T]) Rotate() Matrix[T] {
	h, w := g.Size()
	ng := MakeMatrix[T](w, h)
	for j, row := range g {
		for i := range row {
			ng[i][w-j-1] = row[i]
		}
	}
	return ng
}

// DigitsToMatrix converts a grid of digits into a Matrix[int].
func DigitsToMatrix(digits []string) Matrix[int] {
	h := len(digits)
	if h == 0 {
		return nil
	}
	w := len(digits[0])
	g := MakeMatrix[int](w, h)
	for j, row := range digits {
		for i, d := range row {
			g[j][i] = int(d - '0')
		}
	}
	return g
}