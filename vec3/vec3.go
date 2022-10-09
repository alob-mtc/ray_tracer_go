package vec3

type Vec3 struct {
	e [3]float32
}

func New(e0, e1, e2 float32) *Vec3 {
	return &Vec3{e: [3]float32{e0, e1, e2}}
}

func (v *Vec3) X() float32 {
	return v.e[0]
}
func (v *Vec3) Y() float32 {
	return v.e[1]
}
func (v *Vec3) Z() float32 {
	return v.e[2]
}

func (v *Vec3) R() float32 {
	return v.e[0]
}
func (v *Vec3) G() float32 {
	return v.e[1]
}
func (v *Vec3) B() float32 {
	return v.e[2]
}

func (v *Vec3) Length() float32 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func UnitVector(v *Vec3) *Vec3 {
	return v.Div(v.Length())
}

func Dot(v1, v2 *Vec3) float32 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v1.e[2]*v2.e[2]
}

//=====

func (v Vec3) Add(rhs *Vec3) *Vec3 {
	v.e[0] += rhs.e[0]
	v.e[1] += rhs.e[1]
	v.e[2] += rhs.e[2]
	return &v
}

func (v Vec3) Sub(rhs *Vec3) *Vec3 {
	v.e[0] -= rhs.e[0]
	v.e[1] -= rhs.e[1]
	v.e[2] -= rhs.e[2]
	return &v
}

func (v Vec3) Mul(rhs float32) *Vec3 {
	v.e[0] *= rhs
	v.e[1] *= rhs
	v.e[2] *= rhs
	return &v
}

func (v Vec3) Div(rhs float32) *Vec3 {
	k := 1.0 / rhs
	v.e[0] *= k
	v.e[1] *= k
	v.e[2] *= k
	return &v
}
