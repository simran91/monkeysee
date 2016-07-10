package mods

import "../monkey"
import "image/color"

//
// SwapRGBtoGBR is a mod that swaps the colours around... it's a very simple mod designed
// to show how we loop over the ImageMatrix and read the colour values...
//
func SwapRGBtoGBR(matrix monkey.ImageMatrix) monkey.ImageMatrix {

	width := matrix.GetWidth()
	height := matrix.GetHeight()

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			colour := matrix[x][y].(color.RGBA)
			colour.R, colour.G, colour.B = colour.G, colour.B, colour.R
			matrix[x][y] = colour
		}
	}

	return matrix
}
