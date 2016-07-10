package mods

import "../monkey"
import "image/color"

//
// ApplyFunctionToEveryPixelExample is an example of how to pass a callback function to ApplyFunctionToEveryPixelExample
//
func ApplyFunctionToEveryPixelExample(matrix monkey.ImageMatrix) monkey.ImageMatrix {
	matrix.ApplyFunctionToEveryPixel(rgb2brg)
	return matrix
}

func rgb2brg(im monkey.ImageMatrix, x, y int) color.RGBA {
    c := im[x][y].(color.RGBA)
    c.R, c.G, c.B = c.B, c.R, c.G
	return c
}
