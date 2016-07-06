package mods

import "../mimage"
// import "image/color"

// GaussianBlurConvolution ...
var GaussianBlurConvolution = [][]uint8{
									{1, 2, 1},
									{2, 4, 2},
									{1, 2, 1},
							  }

//
// GaussianBlur performs a gaussian blur...
//
func GaussianBlur(matrix mimage.ImageMatrix) mimage.ImageMatrix {
	newMatrix := matrix.ApplyConvolution(GaussianBlurConvolution)
	return newMatrix
}
