package mods

import "../mimage"

//
// SwapRGBtoGBR is a mod that swaps the colours around... it's a very simple mod designed
// to show how we loop over the ImageMatrix and read the RGBA values...
//
func SwapRGBtoGBR(matrix mimage.ImageMatrix) mimage.ImageMatrix {

	width := len(matrix["r"])
	height := len(matrix["r"][0])

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r := matrix["r"][x][y]
			g := matrix["g"][x][y]
			b := matrix["b"][x][y]
			// a := matrix["a"][x][y]

			// average := uint32((0.21*float64(r) + 0.72*float64(g) + 0.07*float64(b)) / 3)
			// average := (r + g + b) / 3

			matrix["r"][x][y] = g
			matrix["g"][x][y] = b
			matrix["b"][x][y] = r

		}
	}

	return matrix
}
