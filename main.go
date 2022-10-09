package main

import (
	"fmt"
	"github.com/alob-mtc/ray_tracer/ray"
	"github.com/alob-mtc/ray_tracer/vec3"
)

func hitSphere(center vec3.Vec3, radius float32, r *ray.Ray) bool {
	oc := r.Origin().Sub(&center)
	direction := r.Direction()
	a := vec3.Dot(&direction, &direction)
	b := 2 * vec3.Dot(oc, &direction)
	c := vec3.Dot(oc, oc) - radius*radius

	discriminant := b*b - 4*a*c

	return discriminant > 0
}

func color(r *ray.Ray) vec3.Vec3 {
	if hitSphere(*vec3.New(0, 0, -1), 0.5, r) {
		return *vec3.New(1, 0, 0)
	}
	direction := r.Direction()
	unitDirection := vec3.UnitVector(&direction)
	t := 0.5 * (unitDirection.Y() + 1)
	return *vec3.New(1, 1, 1).Mul(1 - t).Add(vec3.New(0.5, 0.7, 1.0).Mul(t))
}

func main() {
	var w int32 = 200
	var h int32 = 100
	var maxValue int32 = 255

	lowerLeftCorner := vec3.New(-2, -1, -1)
	horizontal := vec3.New(4, 0, 0)
	vertical := vec3.New(0, 2, 0)
	origin := vec3.New(0, 0, 0)

	fmt.Printf("P3\n%d %d\n%d\n", w, h, maxValue)

	for j := h - 1; j >= 0; j-- {
		for i := int32(0); i < w; i++ {
			u := float32(i) / float32(w)
			v := float32(j) / float32(h)

			r := ray.New(*origin, *lowerLeftCorner.Add(horizontal.Mul(u)).Add(vertical.Mul(v)))
			col := color(r)

			ir := int32(255.99 * col.R())
			ig := int32(255.99 * col.G())
			ib := int32(255.99 * col.B())

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
