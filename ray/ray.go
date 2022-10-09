package ray

import "github.com/alob-mtc/ray_tracer/vec3"

type Ray struct {
	A vec3.Vec3
	B vec3.Vec3
}

func New(a, b vec3.Vec3) *Ray {
	return &Ray{a, b}
}

func (r *Ray) Origin() vec3.Vec3 {
	return r.A
}

func (r *Ray) Direction() vec3.Vec3 {
	return r.B
}

func (r *Ray) PointOfParameter(t float32) *vec3.Vec3 {
	return r.A.Add(r.B.Mul(t))
}
