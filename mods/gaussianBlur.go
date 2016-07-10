package mods

import "../monkey"

// GaussianBlurConvolution ...
var GaussianBlurConvolution = monkey.ConvolutionMatrix{
	{1, 2, 1},
	{2, 4, 2},
	{1, 2, 1},
}

//
// GaussianBlur performs a gaussian blur...
//
func GaussianBlur(matrix monkey.ImageMatrix) monkey.ImageMatrix {
	newMatrix := matrix.ApplyConvolution(GaussianBlurConvolution)
	return newMatrix
}
