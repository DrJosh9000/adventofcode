package stuff

import "fmt"

var (
	// The reals are a (commutative) division ring (also known as a field).
	_ DivisionRing[Real] = Real(0)
	// The quaternions over the reals are a (non-commutative) division ring.
	_ DivisionRing[Quaternion[Real]] = Quaternion[Real]{}
)

type Ring[T any] interface {
	Add(T) T
	Neg() T
	Mul(T) T
}

type DivisionRing[T any] interface {
	Ring[T]
	Inv() T
}

type Real float64

func (x Real) Add(y Real) Real { return x + y }
func (x Real) Neg() Real       { return -x }
func (x Real) Mul(y Real) Real { return x * y }
func (x Real) Inv() Real       { return 1/x }

type Quaternion[T Ring[T]] [4]T

func (q Quaternion[T]) String() string { return fmt.Sprint(([4]T)(q)) }

func (q Quaternion[T]) Neg() Quaternion[T] {
	return Quaternion[T]{q[0].Neg(), q[1].Neg(), q[2].Neg(), q[3].Neg()}
}

func (q Quaternion[T]) Add(r Quaternion[T]) Quaternion[T] {
	return Quaternion[T]{q[0].Add(r[0]), q[1].Add(r[1]), q[2].Add(r[2]), q[3].Add(r[3])}
}

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

func ExampleQuaternion() {
	var x DivisionRing[Real] = Real(42)
	fmt.Println(x.Mul(69))
	fmt.Println(Quaternion[Real]{1, 2, 3, 4})
	fmt.Println(Quaternion[Real]{1, 4, -3, 0}.Mul(Quaternion[Real]{-1, -1, 2, 7}))
}

type Vec3[T Ring[T]] struct {
	x, y, z T
}

func (v Vec3[T]) Add(w Vec3[T]) Vec3[T] {
	return Vec3[T]{v.x.Add(w.x), v.y.Add(w.y), v.z.Add(w.z)}
}

func (v Vec3[T]) Neg() Vec3[T] {
	return Vec3[T]{v.x.Neg(), v.y.Neg(), v.z.Neg()}
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
