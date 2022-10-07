package main

import "fmt"

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
}
