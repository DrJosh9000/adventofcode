package stuff

import "fmt"

type DivisionRing[T any] interface {
	Add(T) T
	Neg() T
	Mul(T) T
	Inv() T
}

type Quaternion[T DivisionRing[T]] [4]T

func (q Quaternion[T]) String() string { return fmt.Sprint(([4]T)(q)) }

// Inv returns the quaternion inverse.
func (q Quaternion[T]) Inv() Quaternion[T] {
	return Quaternion[T]{q[0], q[1].Neg(), q[2].Neg(), q[3].Neg()}
}

// Mul returns the quaternion product fg.
func (q Quaternion[T]) Mul(r Quaternion[T]) Quaternion[T] {
	return Quaternion[T]{
		q[0].Mul(r[0]).Add(q[1].Mul(r[1]).Neg()).Add(q[2].Mul(r[2]).Neg()).Add(q[3].Mul(r[3]).Neg()),
		q[0].Mul(r[1]).Add(q[1].Mul(r[0])).Add(q[2].Mul(r[3])).Add(q[3].Mul(r[2]).Neg()),
		q[0].Mul(r[2]).Add(q[1].Mul(r[3]).Neg()).Add(q[2].Mul(r[0])).Add(q[3].Mul(r[1])),
		q[0].Mul(r[3]).Add(q[1].Mul(r[2])).Add(q[2].Mul(r[1]).Neg()).Add(q[3].Mul(r[0])),
	}
}

type Real float64

func (x Real) Add(y Real) Real { return x + y }
func (x Real) Neg() Real      { return -x }
func (x Real) Mul(y Real) Real { return x * y }
func (x Real) Inv() Real      { return 1/x }

func ExampleFoo() {
	var x DivisionRing[Real] = Real(42)
	println(x.Mul(69))
	println(Quaternion[Real]{1, 2, 3, 4}.String())
}

type Vec3[T DivisionRing[T]] struct {
	x, y, z T
}

func (v Vec3[T]) Quaternion() Quaternion[T] {
	return Quaternion[T]{1: v.x, 2: v.y, 3: v.z}
}

func (q Quaternion[T]) Vec3() Vec3[T] {
	return Vec3[T]{q[1], q[2], q[3]}
}

func (q Quaternion[T]) Rotate(v Vec3[T]) Vec3[T] {
	return q.Mul(v.Quaternion()).Mul(q.Inv()).Vec3()
}
