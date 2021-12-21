package stuff

// Let's adjoin √2 onto the rationals. This is an algebraic number field.
// Q(√2) = {(a + b√2)/c : a,b,c ∈ ℤ}.

const sqrt2 = 1.414213562373095

var (
	// QR2 is a field (commutative division ring).
	_ DivisionRing[QR2] = QR2{}
	// Quaternions over QR2 are a division ring.
	_ DivisionRing[Quaternion[QR2]] = Quaternion[QR2]{}
)

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

	_ DivisionRing[QR2] = QR2{}
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
