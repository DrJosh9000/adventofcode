package stuff

type Quaternion [4]float64

// Inv returns the quaternion inverse.
func (q Quaternion) Inv() Quaternion {
	return Quaternion{q[0], -q[1], -q[2], -q[3]}
}

// Mul returns the quaternion product qr.
func (q Quaternion) Mul(r Quaternion) Quaternion {
	return Quaternion{
		q[0]*r[0] - q[1]*r[1] - q[2]*r[2] - q[3]*r[3],
		q[0]*r[1] + q[1]*r[0] + q[2]*r[3] - q[3]*r[2],
		q[0]*r[2] - q[1]*r[3] + q[2]*r[0] + q[3]*r[1],
		q[0]*r[3] + q[1]*r[2] - q[2]*r[1] + q[3]*r[0],
	}
}

type Vec struct {
	x, y, z float64
}

func (q Quaternion) Vec() Vec {
	return Vec{q[1], q[2], q[3]}
}

func (q Quaternion) Rotate(v Vec) Vec {
	// Promote v to a quaternion
	p := Quaternion{0, v.x, v.y, v.z}
	// Do the rotation
	r := q.Mul(p).Mul(q.Inv())
	// Return the vector part
	return r.Vec()
}
