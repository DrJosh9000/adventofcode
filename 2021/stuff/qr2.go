package stuff

import "fmt"

// Let's adjoin √2 onto the rationals.
// Q(√2) = {(a + b√2)/c : a,b,c ∈ ℤ}.

const sqrt2 = 1.414213562373095

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

type QR2 struct {
	a, b, c int
}

var (
	QR2MinusOne        = QR2{-1, 0, 1}
	QR2MinusOneOnRoot2 = QR2{0, -1, 2}
	QR2Zero            = QR2{0, 0, 1}
	QR2OneOnRoot2      = QR2{0, 1, 2}
	QR2One             = QR2{1, 0, 1}
)

func (x QR2) Float() float64 {
	return (float64(x.a) + sqrt2*float64(x.b)) / float64(x.c)
}

func (x QR2) Canon() QR2 {
	if x.c < 0 {
		x.a, x.b, x.c = -x.a, -x.b, -x.c
	}
	d := gcd(x.a, gcd(x.b, x.c))
	x.a /= d
	x.b /= d
	x.c /= d
	return x
}

func (x QR2) Neg() QR2 {
	return QR2{-x.a, -x.b, x.c}
}

func (x QR2) Add(y QR2) QR2 {
	//   (a1 + b1√2)/c1 + (a2 + b2√2)/c2
	// = (a1/c1 + a2/c2) + (b1/c1 + b2/c2)√2
	// = ((a1c2 + a2c1) + (b1c2 + b2c1)√2)/c1c2
	return QR2{
		a: x.a*y.c + y.a*x.c,
		b: x.b*y.c + y.b*x.c,
		c: x.c * y.c,
	}.Canon()
}

func (x QR2) Inv() QR2 {
	//   c / (a + b√2)
	// = (c / (a + b√2)) * (a - b√2)/(a - b√2)
	// = c(a - b√2) / (a + b√2)(a - b√2)
	// = (ac - bc√2) / (a² - ab√2 + ab√2 - 2b²)
	// = (ac - bc√2) / (a² - 2b²)
	return QR2{
		a: x.a * x.c,
		b: -x.b * x.c,
		c: x.a*x.a - 2*x.b*x.b,
	}.Canon()
}

func (x QR2) Mul(y QR2) QR2 {
	//   (a1 + b1√2)/c1 * (a2 + b2√2)/c2
	// = (a1 + b1√2)(a2 + b2√2) / c1c2
	// = (a1a2 + a2b1√2 + a1b2√2 + 2b1b2) / c1c2
	// = (a1a2+2b1b2 + (a2b1+a1b2)√2) / c1c2
	return QR2{
		a: x.a*y.a + 2*x.b*y.b,
		b: y.a*x.b + x.a*y.b,
		c: x.c * y.c,
	}.Canon()
}

// Quaternions with this field.
type QR2nion [4]QR2

// Inv returns the quaternion inverse.
func (q QR2nion) Inv() QR2nion {
	return QR2nion{q[0], q[1].Neg(), q[2].Neg(), q[3].Neg()}
}

// Mul returns the quaternion product qr.
func (q QR2nion) Mul(r QR2nion) QR2nion {
	return QR2nion{
		q[0].Mul(r[0]).Add(q[1].Mul(r[1]).Neg()).Add(q[2].Mul(r[2]).Neg()).Add(q[3].Mul(r[3]).Neg()),
		q[0].Mul(r[1]).Add(q[1].Mul(r[0])).Add(q[2].Mul(r[3])).Add(q[3].Mul(r[2]).Neg()),
		q[0].Mul(r[2]).Add(q[1].Mul(r[3]).Neg()).Add(q[2].Mul(r[0])).Add(q[3].Mul(r[1])),
		q[0].Mul(r[3]).Add(q[1].Mul(r[2])).Add(q[2].Mul(r[1]).Neg()).Add(q[3].Mul(r[0])),
	}
}

type QR2Vec struct {
	x, y, z QR2
}

func (v QR2Vec) QR2nion() QR2nion {
	return QR2nion{QR2Zero, v.x, v.y, v.z}
}

func (q QR2nion) Vec() QR2Vec {
	return QR2Vec{q[1], q[2], q[3]}
}

func (q QR2nion) Rotate(v QR2Vec) QR2Vec {
	return q.Mul(v.QR2nion()).Mul(q.Inv()).Vec()
}

type Field[F any] interface {
	Add(F) F
	Neg() F
	Mul(F) F
	Inv() F
}

type footernion[F Field[F]] [4]F

func (f footernion[F]) String() string { return fmt.Sprint(([4]F)(f)) }

// Inv returns the quaternion inverse.
func (f footernion[F]) Inv() footernion[F] {
	return footernion[F]{f[0], f[1].Neg(), f[2].Neg(), f[3].Neg()}
}

// Mul returns the quaternion product qr.
func (f footernion[F]) Mul(g footernion[F]) footernion[F] {
	return footernion[F]{
		f[0].Mul(g[0]).Add(f[1].Mul(g[1]).Neg()).Add(f[2].Mul(g[2]).Neg()).Add(f[3].Mul(g[3]).Neg()),
		f[0].Mul(g[1]).Add(f[1].Mul(g[0])).Add(f[2].Mul(g[3])).Add(f[3].Mul(g[2]).Neg()),
		f[0].Mul(g[2]).Add(f[1].Mul(g[3]).Neg()).Add(f[2].Mul(g[0])).Add(f[3].Mul(g[1])),
		f[0].Mul(g[3]).Add(f[1].Mul(g[2])).Add(f[2].Mul(g[1]).Neg()).Add(f[3].Mul(g[0])),
	}
}

type foo float64

func (f foo) Add(g foo) foo { return f + g }
func (f foo) Neg() foo      { return -f }
func (f foo) Mul(g foo) foo { return f * g }
func (f foo) Inv() foo      { return 1/f }

func exampleFoo() {
	var f Field[foo] = foo(42)
	println(f.Mul(69))
	println(footernion[foo]{1, 2, 3, 4}.String())
}

