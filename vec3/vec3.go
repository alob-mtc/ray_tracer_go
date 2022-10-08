package vec3

type Vec3 struct {
	e [3]float32
}

func New(e0, e1, e2 float32) *Vec3 {
	return &Vec3{e: [3]float32{e0, e1, e2}}
}

func (v Vec3) Add(newV *Vec3) *Vec3 {
	v.e[0] += newV.e[0]
	v.e[1] += newV.e[1]
	v.e[2] += newV.e[2]
	return &v
}
