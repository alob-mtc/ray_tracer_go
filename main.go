package main

import (
	"fmt"
	"github.com/alob-mtc/ray_tracer/vec3"
)

func writePPM(w int32, h int32, max_value int32) {
	fmt.Printf("P3\n%d %d\n%d\n", w, h, max_value)

	for j := h - 1; j >= 0; j-- {
		for i := int32(0); i < w; i++ {
			r := float32(i) / float32(w)
			g := float32(j) / float32(h)
			b := float32(0.2)

			ir := int32(255.99 * r)
			ig := int32(255.99 * g)
			ib := int32(255.99 * b)

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}

func main() {
	var width int32 = 200
	var height int32 = 100
	var maxValue int32 = 255

	writePPM(width, height, maxValue)

	v1 := vec3.New(1, 2, 3)
	v2 := vec3.New(2, 6, 8)

	v3 := v1.Add(v2)

	fmt.Println("Added v1 and v2:", v3)
}
