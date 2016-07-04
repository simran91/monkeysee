package mods

import "../mimage"
import "image/color"

//
// GreyscaleAverageWithTranslusence is a mod that that does a simple average greyscale conversion...
//
func GreyscaleAverageWithTranslusence(matrix mimage.ImageMatrix) mimage.ImageMatrix {

	width := len(matrix)
	height := len(matrix[0])

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			colour := matrix[x][y].(color.RGBA)
			average := (colour.R + colour.G + colour.B) / 3
			colour.R, colour.G, colour.B = average, average, average
			colour.A = colour.A / 2
			matrix[x][y] = colour
		}
	}

	return matrix
}